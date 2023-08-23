package interpreter

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"

	"arc/ast"
	"arc/interpreter/errors"
)

type PresenterHookFn = func(presenter *errors.ErrorPresenter) error

var ErrorManager = &CompilerErrorTracking{
	CurrentSourceFilePath:    "",
	CurrentSourceFileContent: "",
}

type CompilerErrorTracking struct {
	CurrentNode              ast.Node
	CurrentSourceFilePath    string
	CurrentSourceFileContent string
	// CurrentTokenStream       *antlr.CommonTokenStream
	Presenter   *errors.ErrorPresenter
	hooks       []PresenterHookFn
	exitOnError bool
}

func (self *CompilerErrorTracking) SetSource(path string, content string) {
	self.CurrentSourceFileContent = content
	self.CurrentSourceFilePath = path
	if self.CurrentNode != nil {
		self.Presenter = errors.NewErrorPresenter(content, self.CurrentNode.GetRuleRange())
	}
}

func (self *CompilerErrorTracking) SetNode(node ast.Node) {
	self.Presenter = errors.NewErrorPresenter(self.CurrentSourceFileContent, node.GetRuleRange())
	self.CurrentNode = node
}
func (self *CompilerErrorTracking) ShouldExit(value ...bool) {
	val := true
	if len(value) > 0 {
		val = value[0]
	}
	self.exitOnError = val
}

// func (self *CompilerErrorTracking) SetToken(token *lexer.Token) {
// 	self.Presenter = errors.NewErrorPresenter(self.CurrentSourceFileContent, token)
// }

func (self *CompilerErrorTracking) Error(format string, a ...any) {
	log.Helper()
	if self.CurrentNode == nil {
		log.Warnf("CompilerErrorTracking.CurrentNode is nil - falling back to regular log error")
		log.Errorf(format, a...)
		if self.exitOnError {
			os.Exit(1)
		}
		return
	}

	// Log.MarkAsHelperFunc(2)

	self.ErrorAtNode(self.CurrentNode, format, a...)
}

func (self *CompilerErrorTracking) ErrorAtNode(node ast.Node, format string, a ...any) {
	if node == nil {
		panic("node is nil")
	}

	// Errors().AddAtToken(node.GetRuleRange(), format, a...)
	Errors().AddAtNode(node, format, a...)

	if len(self.hooks) > 0 {
		for _, hook := range self.hooks {
			log.Debugf("running presenter hook for error : %s", fmt.Sprintf(format, a...))

			err := hook(Errors())
			if err != nil {
				log.Errorf("Error hook failed: %s", err)
			}
		}
		return
	}

	log.Helper()
	self.Print()
}

func Errors() *errors.ErrorPresenter {
	return ErrorManager.Presenter
}

func (self *CompilerErrorTracking) Print() {
	Errors().Print(self.CurrentSourceFilePath)

	if self.exitOnError {
		os.Exit(1)
	}
}

func (self *CompilerErrorTracking) AddProcessor(cb PresenterHookFn) {
	self.hooks = append(self.hooks, cb)

}

func NewError(format string, a ...any) {
	log.Helper()
	ErrorManager.Error(format, a...)
}

func NewErrorAtNode(node ast.Node, format string, a ...any) {
	log.Helper()
	ErrorManager.SetNode(node)
	ErrorManager.ErrorAtNode(node, format, a...)

}

// func NewErrorAtToken(node *lexer.Token, format string, a ...any) {
// 	log.Helper()
// 	ErrorManager.SetToken(node)
// 	Errors().AddAtNode(node, format, a...)
// 	ErrorManager.Print()
// }
