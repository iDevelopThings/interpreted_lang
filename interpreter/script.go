package interpreter

import (
	"github.com/charmbracelet/log"

	"arc/ast"
)

type SourceFile struct {
	Source string
	Path   string

	// InputStream *antlr.InputStream
	// Stream      *antlr.CommonTokenStream
	// Tree grammar.IProgramContext

	Program *ast.Program

	Env    *Environment
	Logger *log.Logger
}

func (self *SourceFile) GetMainFunc() *ast.FunctionDeclaration {
	if self.Program == nil {
		log.Errorf("Program is nil")
		return nil
	}

	for _, decl := range self.Program.Declarations {
		if funcDecl, ok := decl.(*ast.FunctionDeclaration); ok {
			if funcDecl.Name == "main" {
				return funcDecl
			}
		}
	}

	return nil
}
