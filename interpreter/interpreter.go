package interpreter

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"arc/interpreter/diagnostics"
	"arc/log"

	"arc/ast"
	"arc/http_server"
	"arc/interpreter/config"
	"arc/interpreter/errors"
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
}

func NewInterpreterEngine() *InterpreterEngine {
	env := NewEnvironment()
	RootEnvironment = env

	engine := &InterpreterEngine{
		SourceFiles: make([]*SourceFile, 0),
		Env:         env,
		Evaluator:   NewEvaluator(env),
	}

	log.SetOutput(os.Stderr)

	return engine
}
func NewTestingInterpreterEngine() *InterpreterEngine {
	engine := NewInterpreterEngine()
	engine.IsTesting = true

	Engine = engine

	return engine
}

var Engine = NewInterpreterEngine()

func (self *InterpreterEngine) Load() {
	if config.CliConfig.StdinMode {
		self.LoadSourceFromStdin()
	} else if config.CliConfig.Extra.File != "" {
		self.LoadSource(config.CliConfig.Extra.File)
	}

	switch config.CliConfig.LogOutputMode {
	case config.LogOutputModeStdErr:
		log.SetOutput(os.Stderr)
	case config.LogOutputModeStdOut:
		log.SetOutput(os.Stdout)
	}

	errors.Manager.ConfigureLogger(config.CliConfig.LogOutputMode)
}

func (self *InterpreterEngine) isValidSourcePath(path string) bool {
	if !strings.HasSuffix(path, ".arc") {
		return false
	}

	if filepath.IsAbs(path) {
		if !strings.HasPrefix(path, config.CliConfig.WorkingDirectory) {
			return false
		}
	} else {
		path = filepath.Join(config.CliConfig.WorkingDirectory, path)
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

func (self *InterpreterEngine) LoadSource(path string) *SourceFile {
	source := &SourceFile{
		Path: path,
	}

	src, err := os.ReadFile(source.Path)
	if err != nil {
		panic(err)
	}

	source.Source = string(src)
	source.RelPath, _ = filepath.Rel(config.CliConfig.WorkingDirectory, source.Path)

	return self.FinishLoadingSource(source)
}

func (self *InterpreterEngine) FinishLoadingSource(source *SourceFile) *SourceFile {
	self.setScriptLogger(source)
	defer self.setScriptLogger(nil)

	self.SourceFiles = append(self.SourceFiles, source)

	log.Debugf("Loaded source: %s", source.Path)

	return source
}

func (self *InterpreterEngine) LoadSourceFromString(sourceSrc string) *SourceFile {
	source := &SourceFile{
		Path:    "stdin",
		RelPath: "stdin",
	}

	source.Source = sourceSrc

	self.setScriptLogger(source)
	defer self.setScriptLogger(nil)

	self.SourceFiles = append(self.SourceFiles, source)

	return source
}
func (self *InterpreterEngine) LoadSourceFromStdin() *SourceFile {
	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Error reading from stdin: %v", err)
	}
	source := self.LoadSourceFromString(string(stdin))
	if config.CliConfig.Extra.File != "" {
		source.Path = config.CliConfig.Extra.File
	}

	return source
}

func (self *InterpreterEngine) setScriptLogger(source *SourceFile) {
	if source == nil {
		log.SetPrefix("")
		return
	}

	errors.Manager.SetCurrent(source.Path, source.Source)
	log.SetPrefix(fmt.Sprintf("[%s]", source.RelPath))
}

func (self *InterpreterEngine) parseSource(source *SourceFile) {
	self.setScriptLogger(source)
	defer self.setScriptLogger(nil)

	timer := utilities.NewTimer("Parse & Lex: " + source.Path)
	defer timer.StopAndLog()

	// TODO: Move/cleanup this panic handler
	defer func() {
		if r := recover(); r != nil {
			if !errors.Manager.HandlePanic(r) {
				panic(r)
			}
			errors.TryDumpDiagnostics()
		}

	}()

	l := lexer.NewLexer(source.Source)
	l.SetSource(source.Path)
	p := parser.NewParser(l)

	source.Program = p.Parse()

	if config.CliConfig.PrintAst {
		println(strings.Repeat("-", 80))
		log.Debugf("Printing AST for %s", source.Path)
		source.Print()
		println(strings.Repeat("-", 80))
	}

}

func (self *InterpreterEngine) parseAll() {
	parseTimer := utilities.NewTimer("Parse All SourceFiles")
	defer parseTimer.StopAndLog()

	for _, source := range self.SourceFiles {
		self.parseSource(source)
	}

	// If we have any parsing errors at this point, we should probably just exit...
	// It causes the next steps to fail anyway
	if errors.TryDumpDiagnostics() {
		os.Exit(1)
	}
}

func (self *InterpreterEngine) constructASTs() {
	timer := utilities.NewTimer("Construct All SourceFile ASTs")
	defer timer.StopAndLog()
	defer self.setScriptLogger(nil)

	importedScripts := make(map[string]*SourceFile)

	sourcesToParse := utilities.NewStack[*SourceFile]()

	processScript := func(source *SourceFile) {
		self.setScriptLogger(source)

		importedScripts[source.Path] = source

		if len(source.Program.Imports) == 0 {
			return
		}

		for _, importPath := range source.Program.Imports {
			ip := importPath.Path.Value.(string)
			relativePath := path.Join(path.Dir(source.Path), ip)

			if !self.isValidSourcePath(relativePath) {
				NewDiagnosticAtNode(importPath.Path, diagnostics.InvalidImportFilePath, ip)
				continue
			}

			if _, ok := importedScripts[relativePath]; !ok {
				log.Debugf("Queueing imported source for processing : %s", relativePath)
				sourcesToParse.Push(self.LoadSource(relativePath))
			}
		}
	}

	for _, source := range self.SourceFiles {
		processScript(source)
	}

	for sourcesToParse.Len() > 0 {
		source := sourcesToParse.Pop()
		log.Debugf("Processing queued source : %s", source.Path)
		self.parseSource(source)
		processScript(source)
	}

	for _, source := range self.SourceFiles {
		self.setScriptLogger(source)
	}

}

func (self *InterpreterEngine) linkSource(source *SourceFile) {
	// Essentially bind all declarations to the environment/registry
	// So that we're prepared to evaluate the tree

	self.setScriptLogger(source)
	defer self.setScriptLogger(nil)

	if len(source.Program.Declarations) > 0 {
		for _, declaration := range source.Program.Declarations {
			switch t := declaration.(type) {
			case *ast.HttpBlock:
				self.Evaluator.bindHttpDeclarations(t)
			case *ast.FunctionDeclaration:
				Registry.SetFunction(t)
			case *ast.ObjectDeclaration:
				Registry.SetObject(t)
			case *ast.EnumDeclaration:
				enumRtValue := Registry.SetEnum(t)
				self.Env.SetVar(t.Name.Name, enumRtValue)
				self.Evaluator.evalEnumDeclaration(t, enumRtValue)
			}
		}
	}

}

func (self *InterpreterEngine) typeCheckSources() {
	if self.DisableTypeChecker {
		log.Debugf("TypeChecker disabled, skipping type checking")
		return
	}

	timer := utilities.NewTimer("TypeCheck All SourceFiles")
	defer timer.StopAndLog()

	TypeChecker.IsTypeChecking(true)
	defer TypeChecker.IsTypeChecking(false)

	defer self.setScriptLogger(nil)

	for _, source := range self.SourceFiles {
		self.setScriptLogger(source)

		TypeChecker.TypeCheckTree(source.Program, self.Env)
	}
}

func (self *InterpreterEngine) evaluateAll() {
	timer := utilities.NewTimer("Evaluate All SourceFiles")
	defer timer.StopAndLog()
	defer self.setScriptLogger(nil)

	for _, source := range self.SourceFiles {
		self.setScriptLogger(source)

		if source.Env == nil {
			source.Env = self.Env.NewChild()
		}

		self.Evaluator.Eval(source.Program)
	}
}

func (self *InterpreterEngine) ProcessScripts() {
	RegisterRuntimeFunctions(self.Env)

	self.parseAll()
	self.constructASTs()

	for _, source := range self.SourceFiles {
		self.linkSource(source)
	}

	self.typeCheckSources()
}

func (self *InterpreterEngine) Lint() {
	errors.SetStrategy(errors.AccumulateAll)

	runTimer := utilities.NewTimer("Run All SourceFiles")
	defer runTimer.StopAndLog()

	defer errors.TryDumpDiagnostics()

	self.ProcessScripts()
}

func (self *InterpreterEngine) Run() {
	runTimer := utilities.NewTimer("Run All SourceFiles")
	defer runTimer.StopAndLog()

	self.ProcessScripts()

	self.evaluateAll()
	self.runMainAndServer()
}

func (self *InterpreterEngine) runMainAndServer() {
	mainFn := Registry.LookupFunction("main")
	if mainFn != nil {
		self.Evaluator.ForceExecuteFunction(mainFn)
	} else {
		if !self.IsTesting {
			panic("No main function")
		}
	}

	if !self.IsTesting {
		if len(Registry.HttpEnv.Routes) > 0 {
			self.runHttpServer()
		}
	}

}

func (self *InterpreterEngine) runHttpServer() {
	conf := config.ProjectConfig.HttpServer

	router := http_server.GetRouter()

	addr := conf.Address.Value
	port := conf.Port.Value
	host := addr + ":" + strconv.Itoa(port)

	log.Printf("Http Server running http://%s", host)

	server := &http.Server{
		Addr:              host,
		Handler:           router,
		ReadHeaderTimeout: time.Duration(conf.ReadHeaderTimeout.Value) * time.Millisecond,
		WriteTimeout:      time.Duration(conf.WriteTimeout.Value) * time.Millisecond,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
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

// func (self *InterpreterEngine) RunScript(source *SourceFile) {
// 	lexer := grammar.NewSimpleLangLexer(source.InputStream)
// 	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
// 	parser := grammar.NewSimpleLangParser(stream)
//
// 	// This is the tree for file a for example
// 	tree := parser.Program()
//
// }
