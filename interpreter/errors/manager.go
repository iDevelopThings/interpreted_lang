package errors

import (
	"os"

	"github.com/charmbracelet/log"

	"arc/interpreter/config"
)

type ErrorHandlingStrategy string

const (
	ExitOnError     ErrorHandlingStrategy = "exit"
	ContinueOnError ErrorHandlingStrategy = "continue"
)

type FileDiagnosticData struct {
	Diagnostics []CodeDiagnostic
	Path        string
	Source      string
	Presenter   *DiagnosticPresenter
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

func (self *DiagnosticManagerInstance) SetCurrent(path string, source string) {
	_, ok := self.fileDiagnostics[path]
	if !ok {
		self.fileDiagnostics[path] = &FileDiagnosticData{
			Path:        path,
			Source:      source,
			Presenter:   NewPresenter(path, source),
			Diagnostics: []CodeDiagnostic{},
		}
	}

	self.currentPath = path
	self.currentSource = source
}

func (self *DiagnosticManagerInstance) onNew(e *CodeDiagnostic) {
	PresenterLogger.Helper()

	fileData, ok := self.fileDiagnostics[self.currentPath]
	if !ok {
		log.Warnf("DiagnosticManagerInstance.onNew: fileData is nil")
		return
	}

	fileData.Diagnostics = append(fileData.Diagnostics, *e)

	switch self.strategy {
	case ExitOnError:
		fileData.Presenter.Diagnostics = append(fileData.Presenter.Diagnostics, *e)
		fileData.Presenter.Print(self.outputFormat)
		os.Exit(1)

	case ContinueOnError:
		fileData.Presenter.Diagnostics = append(fileData.Presenter.Diagnostics, *e)
		fileData.Presenter.Print(self.outputFormat)
	}
}

func (self *DiagnosticManagerInstance) AddPresentableError(err PresentableError) {
	PresenterLogger.Helper()

	e := NewCodeErrorInRange(err.GetPosition())
	e.Severity = err.GetSeverity()
	e.AddMessage(err.GetMessage())

	self.onNew(e)
}
func (self *DiagnosticManagerInstance) AddNodeDiagnostic(err PresentableNodeError) {
	PresenterLogger.Helper()

	e := NewCodeErrorAtNode(err.GetNode())
	e.Severity = err.GetSeverity()
	e.AddMessage(err.GetMessage())

	self.onNew(e)
}

func (self *DiagnosticManagerInstance) HandlePanic(err any) bool {
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
