package interpreter

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path"
	"strconv"
	"time"

	"github.com/antlr4-go/antlr/v4"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"

	"interpreted_lang/grammar"
	"interpreted_lang/http_server"
	"interpreted_lang/utilities"
)

type InterpreterEngine struct {
	Scripts   []*Script
	Env       *Environment
	Evaluator *Evaluator

	IsTesting bool
}

var globalLogger *log.Logger

func NewInterpreterEngine() *InterpreterEngine {
	env := NewEnvironment()
	engine := &InterpreterEngine{
		Scripts:   make([]*Script, 0),
		Env:       env,
		Evaluator: NewEvaluator(env),
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
	return engine
}

var Engine = NewInterpreterEngine()

func (self *InterpreterEngine) LoadScript(path string) *Script {
	script := &Script{
		Path: path,
	}

	src, err := os.ReadFile(script.Path)
	if err != nil {
		panic(err)
	}

	script.Source = string(src)
	script.InputStream = antlr.NewInputStream(script.Source)

	self.setScriptLogger(script)
	defer self.setScriptLogger(nil)

	self.Scripts = append(self.Scripts, script)

	fmt.Printf("Loaded script: %s\n", path)

	return script
}

func (self *InterpreterEngine) LoadScriptFromString(scriptSrc string) {
	script := &Script{
		Path: "stdin",
	}

	script.Source = scriptSrc
	script.InputStream = antlr.NewInputStream(script.Source)

	self.setScriptLogger(script)
	defer self.setScriptLogger(nil)

	self.Scripts = append(self.Scripts, script)
}

func (self *InterpreterEngine) setScriptLogger(script *Script) {
	if script != nil {
		if script.Logger == nil {
			script.Logger = log.New(os.Stderr)
			script.Logger.SetReportTimestamp(false)
			script.Logger.SetLevel(log.DebugLevel)
			script.Logger.SetReportCaller(true)
			script.Logger.SetPrefix(fmt.Sprintf("[%s]", script.Path))
		}

		ErrorManager.SetSource(
			script.Path,
			script.Source,
			script.Stream,
		)

		log.SetDefault(script.Logger)
	} else {
		log.SetDefault(globalLogger)
	}
}

func (self *InterpreterEngine) parseAll() {
	parseTimer := utilities.NewTimer("Parse All Scripts")
	defer parseTimer.StopAndLog()

	for _, script := range self.Scripts {
		self.parseScript(script)
	}
}

func (self *InterpreterEngine) parseScript(script *Script) {
	self.setScriptLogger(script)
	defer self.setScriptLogger(nil)

	lexer := grammar.NewSimpleLangLexer(script.InputStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	script.Stream = stream
	parser := grammar.NewSimpleLangParser(stream)
	script.Tree = parser.Program()
}

func (self *InterpreterEngine) constructASTs() {
	timer := utilities.NewTimer("Construct All Script ASTs")
	defer timer.StopAndLog()

	importedScripts := make(map[string]*Script)

	scriptsToParse := utilities.NewStack[*Script]()

	processScript := func(script *Script) {
		self.setScriptLogger(script)

		importedScripts[script.Path] = script

		mapper, program := NewAstMapper(script.Tree)

		script.Mapper = mapper
		script.Program = program

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

	for _, script := range self.Scripts {
		processScript(script)
	}

	for scriptsToParse.Len() > 0 {
		script := scriptsToParse.Pop()
		log.Debugf("Processing queued script : %s", script.Path)
		self.parseScript(script)
		processScript(script)
	}

	for _, script := range self.Scripts {
		self.setScriptLogger(script)
		script.Mapper.VisitProgram(script.Tree.(*grammar.ProgramContext))
	}

	defer self.setScriptLogger(nil)
}

func (self *InterpreterEngine) link() {
	for _, script := range self.Scripts {
		self.linkScript(script)
	}
}

func (self *InterpreterEngine) linkScript(script *Script) {
	// Populate the GlobalSymbols table based on global declarations in the tree
	// For example, if a tree declares a global function, add it to the symbol table
	self.setScriptLogger(script)
	defer self.setScriptLogger(nil)

	if len(script.Mapper.Objects) > 0 {
		for _, obj := range script.Mapper.Objects {
			self.Env.SetObject(obj)
		}
	}

	if len(script.Mapper.Functions) > 0 {
		for _, fn := range script.Mapper.Functions {
			self.Env.SetFunction(fn)
		}
	}
}

func (self *InterpreterEngine) typeCheck() {
	timer := utilities.NewTimer("TypeCheck All Scripts")
	defer timer.StopAndLog()

	for _, script := range self.Scripts {
		self.typeCheckScript(script)
	}
}

func (self *InterpreterEngine) typeCheckScript(script *Script) {
	// Populate the GlobalSymbols table based on global declarations in the tree
	// For example, if a tree declares a global function, add it to the symbol table
	self.setScriptLogger(script)
	defer self.setScriptLogger(nil)

	checker := NewTypeCheckingVisitor(script.Program, self.Env)
	if checker != nil {

	}
}

func (self *InterpreterEngine) evaluateAll() {
	timer := utilities.NewTimer("Evaluate All Scripts")
	defer timer.StopAndLog()

	for _, script := range self.Scripts {
		self.evaluateScript(script)
	}
}

func (self *InterpreterEngine) evaluateScript(script *Script) {
	// Evaluate the tree using the GlobalSymbols table to resolve any references
	self.setScriptLogger(script)
	defer self.setScriptLogger(nil)

	env := self.Env.NewChild()
	script.Env = env

	self.Evaluator.Eval(script.Program)

	for _, function := range script.Mapper.Functions {
		if function.Name == "init" {
			self.Evaluator.ExecuteFunction(function)
			break
		}
	}

}

func (self *InterpreterEngine) prepareToEvaluate() {
	RegisterRuntimeFunctions(self.Env)

	self.parseAll()
	self.constructASTs()
	self.link()
	self.typeCheck()
}

func (self *InterpreterEngine) Run() {
	runTimer := utilities.NewTimer("Run All Scripts")
	defer runTimer.StopAndLog()

	self.prepareToEvaluate()

	self.evaluateAll()

	// wg := &sync.WaitGroup{}

	self.runMainAndServer()
}

func (self *InterpreterEngine) runMainAndServer() {
	mainFn := self.Env.LookupFunction("main")
	if mainFn != nil {
		self.Evaluator.ExecuteFunction(mainFn)
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

	fmt.Println("Done")
}

// func (self *InterpreterEngine) RunScript(script *Script) {
// 	lexer := grammar.NewSimpleLangLexer(script.InputStream)
// 	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
// 	parser := grammar.NewSimpleLangParser(stream)
//
// 	// This is the tree for file a for example
// 	tree := parser.Program()
//
// }
