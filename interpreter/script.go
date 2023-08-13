package interpreter

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/charmbracelet/log"

	"interpreted_lang/ast"
	"interpreted_lang/grammar"
)

type Script struct {
	Source string
	Path   string

	InputStream *antlr.InputStream
	Stream      *antlr.CommonTokenStream

	Tree grammar.IProgramContext

	Mapper  *AstMapper
	Program *ast.Program

	Env    *Environment
	Logger *log.Logger
}

func (self *Script) GetMainFunc() *ast.FunctionDeclaration {
	if self.Program == nil {
		log.Errorf("Program is nil")
		return nil
	}

	for _, statement := range self.Program.Statements {
		if funcDecl, ok := statement.(*ast.FunctionDeclaration); ok {
			if funcDecl.Name == "main" {
				return funcDecl
			}
		}
	}

	return nil
}
