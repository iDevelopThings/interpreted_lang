package errors

import (
	"os"

	"arc/log"

	"arc/interpreter/config"
)

type ErrorHandlingStrategy string

const (
	ExitOnError     ErrorHandlingStrategy = "exit"
	ContinueOnError ErrorHandlingStrategy = "continue"
	AccumulateAll   ErrorHandlingStrategy = "accumulate"
)

type FileDiagnosticData struct {
	// Diagnostics []CodeDiagnostic
	Path      string
	Source    string
	Presenter *DiagnosticPresenter
}

type DiagnosticManagerInstance struct {
	strategy     ErrorHandlingStrategy
	outputFormat config.OutputFormat

	// The engine will tell the manager that we're processing the path/source
	// Any errors that occur during the parse/lex phase will then be assocated with
	// the current path/source
	//
	// **** These should not be used outside of the parse/lex phase ****
	//
	currentPath   string
	currentSource string

	fileDiagnostics map[string]*FileDiagnosticData
}

var Manager = NewDiagnosticManager()

func NewDiagnosticManager() *DiagnosticManagerInstance {
	inst := &DiagnosticManagerInstance{
		strategy:        ExitOnError,
		fileDiagnostics: map[string]*FileDiagnosticData{},
	}

	return inst
}

func (self *DiagnosticManagerInstance) ConfigureLogger(mode config.LogOutputMode) {
	switch mode {
	case config.LogOutputModeStdErr:
		PresenterLogger.SetOutput(os.Stderr)
	case config.LogOutputModeStdOut:
		PresenterLogger.SetOutput(os.Stdout)
	}
}

func (self *DiagnosticManagerInstance) SetCurrent(path string, source string) {
	_, ok := self.fileDiagnostics[path]
	if !ok {
		self.fileDiagnostics[path] = &FileDiagnosticData{
			Path:      path,
			Source:    source,
			Presenter: NewPresenter(path, source),
		}
	}

	self.currentPath = path
	self.currentSource = source
}

func (self *DiagnosticManagerInstance) pushDiagnostic(e CodeDiagnostic) *FileDiagnosticData {
	PresenterLogger.Helper()

	fileData, ok := self.fileDiagnostics[self.currentPath]
	if !ok {
		log.Warnf("DiagnosticManagerInstance.onNew: fileData is nil")
		return nil
	}

	// fileData.Diagnostics = append(fileData.Diagnostics, e)
	fileData.Presenter.Diagnostics = append(fileData.Presenter.Diagnostics, e)

	return fileData
}

func (self *DiagnosticManagerInstance) onNew(e CodeDiagnostic) {
	PresenterLogger.Helper()

	fileData := self.pushDiagnostic(e)
	shouldIgnoreStrategy := e.Severity == WarningDiagnostic
	self.printDiagnostics(fileData, shouldIgnoreStrategy)
}

func (self *DiagnosticManagerInstance) printDiagnostics(fileData *FileDiagnosticData, shouldIgnoreStrategy ...bool) bool {
	PresenterLogger.Helper()

	ignoreStrategy := false
	if len(shouldIgnoreStrategy) > 0 && shouldIgnoreStrategy[0] {
		ignoreStrategy = true
	}

	if self.strategy == AccumulateAll && !ignoreStrategy {
		return false
	}

	if len(fileData.Presenter.Diagnostics) == 0 {
		return false
	}

	fileData.Presenter.Print(self.outputFormat)

	if ignoreStrategy {
		return false
	}

	if self.strategy == ExitOnError {
		os.Exit(1)
	}

	return true
}

func (self *DiagnosticManagerInstance) AddPresentableError(err PresentableError) {
	PresenterLogger.Helper()

	e := NewCodeErrorInRange(err.GetPosition())
	e.Severity = err.GetSeverity()
	e.AddMessage(err.GetMessage())

	if err, ok := err.(ErrorWithCallerInfo); ok {
		e.SetCallerInfo(err.GetCallerInfo())
	} else {
		e.SetCallerInfo(PresenterLogger.GetCallerInfo())
	}

	self.onNew(*e)
}
func (self *DiagnosticManagerInstance) AddNodeDiagnostic(err PresentableNodeError) {
	PresenterLogger.Helper()

	e := NewCodeErrorAtNode(err.GetNode())
	e.Severity = err.GetSeverity()
	e.Code = err.GetCode()
	e.AddMessage(err.GetMessage())

	if err, ok := err.(ErrorWithCallerInfo); ok {
		e.SetCallerInfo(err.GetCallerInfo())
	} else {
		e.SetCallerInfo(PresenterLogger.GetCallerInfo())
	}

	self.onNew(*e)
}

func (self *DiagnosticManagerInstance) pushTempBuilder(builder *TempDiagnosticBuilder) {
	PresenterLogger.Helper()

	var fileData *FileDiagnosticData
	for _, diagnostic := range builder.Diagnostics {
		fileData = self.pushDiagnostic(diagnostic)
	}

	self.printDiagnostics(fileData, false)
}

func (self *DiagnosticManagerInstance) HandlePanic(err any) bool {
	log.Helper()
	switch e := err.(type) {
	case PresentableError:
		self.AddPresentableError(e)
		return true
	case PresentableNodeError:
		self.AddNodeDiagnostic(e)
		return true
	default:
		return false
	}

}

func SetStrategy(strategy ErrorHandlingStrategy) {
	Manager.strategy = strategy
}

func SetFormat(format config.OutputFormat) {
	Manager.outputFormat = format
}

func TryDumpDiagnostics() bool {
	PresenterLogger.Helper()

	numPrinted := 0
	if Manager.strategy == AccumulateAll {
		for _, fileData := range Manager.fileDiagnostics {
			if len(fileData.Presenter.Diagnostics) == 0 {
				continue
			}
			Manager.printDiagnostics(fileData, true)
			numPrinted++
		}
	}

	return numPrinted > 0
}
