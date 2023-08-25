package interpreter

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"path"
	"strconv"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"

	"arc/ast"
	"arc/http_server"
	"arc/lexer"
	"arc/parser"
	"arc/utilities"
)

type InterpreterEngine struct {
	SourceFiles []*SourceFile
	Env         *Environment
	Evaluator   *Evaluator

	IsTesting          bool
	DisableTypeChecker bool
	loggingEnabled     bool
	loggingWriter      io.Writer
}

var globalLogger *log.Logger

func NewInterpreterEngine() *InterpreterEngine {
	env := NewEnvironment()
	engine := &InterpreterEngine{
		loggingEnabled: true,

		SourceFiles: make([]*SourceFile, 0),
		Env:         env,
		Evaluator:   NewEvaluator(env),
	}

	logger := log.Default()
	logger.SetReportTimestamp(false)
	logger.SetReportCaller(true)
	logger.SetLevel(log.DebugLevel)

	baseStyle := lipgloss.NewStyle().
		Padding(0, 1, 0, 1).
		Bold(true).
		MaxWidth(7).
		Foreground(lipgloss.Color("#fafafa"))

	log.DebugLevelStyle = baseStyle.Copy().
		SetString("DEBUG").
		Background(lipgloss.Color("#6366f1"))

	log.WarnLevelStyle = baseStyle.Copy().
		SetString("WARN").
		Padding(0, 1, 0, 2).
		Background(lipgloss.Color("#facc15")).
		Foreground(lipgloss.Color("#000000"))

	log.ErrorLevelStyle = baseStyle.Copy().
		SetString("ERROR").
		Background(lipgloss.Color("#f87171"))

	log.FatalLevelStyle = baseStyle.Copy().
		SetString("FATAL").
		Background(lipgloss.Color("#dc2626"))

	logTmp := *logger
	globalLogger = &logTmp

	return engine
}
func NewTestingInterpreterEngine() *InterpreterEngine {
	engine := NewInterpreterEngine()
	engine.IsTesting = true

	Engine = engine

	return engine
}

var Engine = NewInterpreterEngine()

func (self *InterpreterEngine) FinishLoadingScript(script *SourceFile) *SourceFile {
	// script.InputStream = antlr.NewInputStream(script.Source)

	self.setScriptLogger(script)
	defer self.setScriptLogger(nil)

	self.SourceFiles = append(self.SourceFiles, script)

	log.Debugf("Loaded script: %s\n", script.Path)

	return script
}

func (self *InterpreterEngine) LoadScript(path string) *SourceFile {
	script := &SourceFile{
		Path: path,
	}

	src, err := os.ReadFile(script.Path)
	if err != nil {
		panic(err)
	}

	script.Source = string(src)

	return self.FinishLoadingScript(script)
}

func (self *InterpreterEngine) LoadScriptFromString(scriptSrc string) {
	script := &SourceFile{
		Path: "stdin",
	}

	script.Source = scriptSrc
	// script.InputStream = antlr.NewInputStream(script.Source)

	self.setScriptLogger(script)
	defer self.setScriptLogger(nil)

	self.SourceFiles = append(self.SourceFiles, script)
}

func (self *InterpreterEngine) setScriptLogger(script *SourceFile) {
	if script != nil {
		if script.Logger == nil {
			script.Logger = log.New(os.Stderr)
			script.Logger.SetReportTimestamp(false)
			if self.loggingEnabled {
				script.Logger.SetLevel(log.DebugLevel)
			} else {
				script.Logger.SetLevel(log.ErrorLevel)
				script.Logger.SetOutput(self.loggingWriter)
			}
			script.Logger.SetReportCaller(true)
			script.Logger.SetPrefix(fmt.Sprintf("[%s]", script.Path))
		}

		ErrorManager.SetSource(
			script.Path,
			script.Source,
		)

		log.SetDefault(script.Logger)
	} else {
		log.SetDefault(globalLogger)
	}
}

func (self *InterpreterEngine) parseAll() {
	parseTimer := utilities.NewTimer("Parse All SourceFiles")
	defer parseTimer.StopAndLog()

	for _, script := range self.SourceFiles {
		self.parseScript(script)
	}
}

func (self *InterpreterEngine) parseScript(script *SourceFile) {
	self.setScriptLogger(script)
	defer self.setScriptLogger(nil)

	timer := utilities.NewTimer("Parse Source: " + script.Path)
	defer timer.StopAndLog()

	// lexer := grammar.NewSimpleLangLexer(script.InputStream)
	// stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// script.Stream = stream
	// parser := grammar.NewSimpleLangParser(stream)
	// script.Tree = parser.Program()

	l := lexer.NewLexer(script.Source)
	l.SetSource(script.Path)
	p := parser.NewParser(l)

	script.Program = p.Parse()
}

func (self *InterpreterEngine) constructASTs() {
	timer := utilities.NewTimer("Construct All SourceFile ASTs")
	defer timer.StopAndLog()

	importedScripts := make(map[string]*SourceFile)

	scriptsToParse := utilities.NewStack[*SourceFile]()

	processScript := func(script *SourceFile) {
		self.setScriptLogger(script)

		importedScripts[script.Path] = script

		if len(script.Program.Imports) > 0 {
			for _, importPath := range script.Program.Imports {
				ip := importPath.Path.Value.(string)
				relativePath := path.Join(path.Dir(script.Path), ip)
				if _, ok := importedScripts[relativePath]; !ok {
					log.Debugf("Queueing imported script for processing : %s", relativePath)
					scriptsToParse.Push(self.LoadScript(relativePath))
				}
			}
		}
	}

	for _, script := range self.SourceFiles {
		processScript(script)
	}

	for scriptsToParse.Len() > 0 {
		script := scriptsToParse.Pop()
		log.Debugf("Processing queued script : %s", script.Path)
		self.parseScript(script)
		processScript(script)
	}

	for _, script := range self.SourceFiles {
		self.setScriptLogger(script)
		// script.Mapper.VisitProgram(script.Tree.(*grammar.ProgramContext))
	}

	defer self.setScriptLogger(nil)
}

func (self *InterpreterEngine) link() {
	for _, script := range self.SourceFiles {
		self.linkScript(script)
	}
}

func (self *InterpreterEngine) linkScript(script *SourceFile) {
	// Populate the GlobalSymbols table based on global declarations in the tree
	// For example, if a tree declares a global function, add it to the symbol table
	self.setScriptLogger(script)
	defer self.setScriptLogger(nil)

	if len(script.Program.Declarations) > 0 {
		for _, declaration := range script.Program.Declarations {
			switch t := declaration.(type) {
			case *ast.FunctionDeclaration:
				self.Env.SetFunction(t)
			case *ast.ObjectDeclaration:
				self.Env.SetObject(t)
			case *ast.EnumDeclaration:
				self.Env.SetEnum(t, self.Evaluator)

			}
		}
	}

}

func (self *InterpreterEngine) typeCheck() {
	if self.DisableTypeChecker {
		log.Debugf("TypeChecker disabled, skipping type checking")
		return
	}

	timer := utilities.NewTimer("TypeCheck All SourceFiles")
	defer timer.StopAndLog()

	TypeChecker.IsTypeChecking(true)
	defer TypeChecker.IsTypeChecking(false)

	for _, script := range self.SourceFiles {
		self.typeCheckScript(script)
	}
}

func (self *InterpreterEngine) typeCheckScript(script *SourceFile) {
	// Populate the GlobalSymbols table based on global declarations in the tree
	// For example, if a tree declares a global function, add it to the symbol table
	self.setScriptLogger(script)
	defer self.setScriptLogger(nil)

	checker := NewTypeCheckingVisitor(script.Program, self.Env)
	if checker != nil {

	}
}

func (self *InterpreterEngine) evaluateAll() {
	timer := utilities.NewTimer("Evaluate All SourceFiles")
	defer timer.StopAndLog()

	for _, script := range self.SourceFiles {
		self.evaluateScript(script)
	}
}

func (self *InterpreterEngine) evaluateScript(script *SourceFile) {
	// Evaluate the tree using the GlobalSymbols table to resolve any references
	self.setScriptLogger(script)
	defer self.setScriptLogger(nil)

	env := self.Env.NewChild()
	script.Env = env

	self.Evaluator.Eval(script.Program)

	// if mainFunc := script.GetMainFunc(); mainFunc != nil {
	// 	self.Evaluator.ForceExecuteFunction(mainFunc)
	// }

}

func (self *InterpreterEngine) ProcessScripts() {
	RegisterRuntimeFunctions(self.Env)

	self.parseAll()
	self.constructASTs()
	self.link()
	self.typeCheck()
}

func (self *InterpreterEngine) Run() {
	runTimer := utilities.NewTimer("Run All SourceFiles")
	defer runTimer.StopAndLog()

	self.ProcessScripts()

	self.evaluateAll()

	// wg := &sync.WaitGroup{}

	w := utilities.NewIndentWriter(os.Stdout, " ")
	self.SourceFiles[0].Program.PrintTree(w.(*utilities.IndentWriter))

	self.runMainAndServer()
}

func (self *InterpreterEngine) runMainAndServer() {
	mainFn := self.Env.LookupFunction("main")
	if mainFn != nil {
		self.Evaluator.ForceExecuteFunction(mainFn)
	} else {
		if !self.IsTesting {
			panic("No main function")
		}
	}

	if !self.IsTesting {
		if len(self.Env.HttpEnv.Routes) > 0 {
			router := http_server.GetRouter()
			// wg.Add(1)

			strPort := strconv.Itoa(router.Options.Port)
			log.Printf("Http Server running http://127.0.0.1:%v\n", router.Options.Port)

			server := &http.Server{Addr: ":" + strPort, Handler: router}

			go func() {
				if err := server.ListenAndServe(); err != nil {
					// handle err
					log.Errorf("Error starting server: %v", err)
				}
			}()

			// Setting up signal capturing
			stop := make(chan os.Signal, 1)
			signal.Notify(stop, os.Interrupt)

			// Waiting for SIGINT (kill -2)
			<-stop

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := server.Shutdown(ctx); err != nil {
				// handle err
				log.Errorf("Error shutting down server: %v", err)
			}
		}

		// wg.Wait()
	}

}

func (self *InterpreterEngine) DisableLogging(w io.Writer) {
	globalLogger.SetLevel(log.FatalLevel)
	globalLogger.SetOutput(w)
	self.loggingWriter = w
	self.loggingEnabled = false
}

// func (self *InterpreterEngine) RunScript(script *SourceFile) {
// 	lexer := grammar.NewSimpleLangLexer(script.InputStream)
// 	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
// 	parser := grammar.NewSimpleLangParser(stream)
//
// 	// This is the tree for file a for example
// 	tree := parser.Program()
//
// }
