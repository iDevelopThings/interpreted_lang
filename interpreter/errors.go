package interpreter

import (
	"fmt"
	"os"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/charmbracelet/log"

	"interpreted_lang/ast"
	"interpreted_lang/interpreter/errors"
)

var ErrorManager = &CompilerErrorTracking{
	CurrentSourceFilePath:    "",
	CurrentSourceFileContent: "",
}

type CompilerErrorTracking struct {
	CurrentNode              ast.Node
	CurrentSourceFilePath    string
	CurrentSourceFileContent string
	CurrentTokenStream       *antlr.CommonTokenStream
	Presenter                *errors.ErrorPresenter
}

func (self *CompilerErrorTracking) SetSource(path string, content string, stream *antlr.CommonTokenStream) {
	self.CurrentSourceFileContent = content
	self.CurrentSourceFilePath = path
	self.CurrentTokenStream = stream
	if self.CurrentNode != nil {
		self.Presenter = errors.NewErrorPresenter(content, self.CurrentNode.GetRule())
	}
}

func (self *CompilerErrorTracking) SetNode(node ast.Node) {
	self.Presenter = errors.NewErrorPresenter(self.CurrentSourceFileContent, node.GetRule())
	self.CurrentNode = node
}
func (self *CompilerErrorTracking) SetRule(rule antlr.ParserRuleContext) {
	self.Presenter = errors.NewErrorPresenter(self.CurrentSourceFileContent, rule)
}

func (self *CompilerErrorTracking) Error(format string, a ...any) {
	if self.CurrentNode == nil {
		log.Warnf("CompilerErrorTracking.CurrentNode is nil - falling back to regular log error")
		log.Errorf(format, a...)
		os.Exit(1)
		return
	}

	// Log.MarkAsHelperFunc(2)

	self.ErrorAtNode(self.CurrentNode, format, a...)
}

func (self *CompilerErrorTracking) ErrorAtNode(node ast.Node, format string, a ...any) {
	if node == nil {
		panic("node is nil")
	}

	Errors().AddAtToken(node.GetRule(), format, a...)

	log.Helper()
	self.Print()
}

func Errors() *errors.ErrorPresenter {
	return ErrorManager.Presenter
}

func (self *CompilerErrorTracking) Print() {
	Errors().Print(self.CurrentSourceFilePath)

	// callerInfo := log.CallerInfo(2)
	log.Helper()
	log.Debugf("From: %s\n", self.CurrentSourceFilePath)
	fmt.Println(strings.Repeat("-", 80))

	// fmt.Printf("From: %s:%d\n", callerInfo.File, callerInfo.Line)
	// fmt.Println(strings.Repeat("-", 80))
	// fmt.Println("")

	os.Exit(1)
}

func NewError(format string, a ...any) {
	ErrorManager.Error(format, a...)
}

func NewErrorAtNode(node ast.Node, format string, a ...any) {
	log.Helper()
	ErrorManager.SetNode(node)
	ErrorManager.ErrorAtNode(node, format, a...)

}

func NewErrorAtToken(node antlr.ParserRuleContext, format string, a ...any) {
	log.Helper()
	ErrorManager.SetRule(node)
	Errors().AddAtToken(node, format, a...)
	ErrorManager.Print()
}
