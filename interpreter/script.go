package interpreter

import (
	"os"

	"arc/ast"
	"arc/log"
	"arc/utilities"
)

type SourceFile struct {
	Source string

	Path    string
	RelPath string

	Program *ast.Program

	Env *Environment
}

func (self *SourceFile) GetFunc(name string) *ast.FunctionDeclaration {
	if self.Program == nil {
		log.Errorf("Program is nil")
		return nil
	}

	for _, decl := range self.Program.Declarations {
		if funcDecl, ok := decl.(*ast.FunctionDeclaration); ok {
			if funcDecl.Name == name {
				return funcDecl
			}
		}
	}

	return nil
}
func (self *SourceFile) GetMainFunc() *ast.FunctionDeclaration {
	return self.GetFunc("main")
}

func (self *SourceFile) Print() {
	w := utilities.NewIndentWriter(os.Stdout, "  ")
	self.Program.PrintTree(w.(*utilities.IndentWriter))
}
