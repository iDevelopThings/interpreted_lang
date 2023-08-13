package interpreter

import (
	"interpreted_lang/ast"
)

type TypeScope struct {
	Types []map[string]ast.Type
}

func NewTypeScope() *TypeScope {
	return &TypeScope{
		Types: []map[string]ast.Type{},
	}
}

func (self *TypeScope) Push() {
	self.Types = append(self.Types, map[string]ast.Type{})
}

func (self *TypeScope) Pop() {
	self.Types = self.Types[:len(self.Types)-1]
}

func (self *TypeScope) Insert(name string, typ ast.Type) {
	self.Types[len(self.Types)-1][name] = typ
}

func (self *TypeScope) Lookup(name string) ast.Type {
	for i := len(self.Types) - 1; i >= 0; i-- {
		if typ, ok := self.Types[i][name]; ok {
			return typ
		}
	}

	return nil
}

func (self *TypeScope) IsDefined(name string) bool {
	return self.Lookup(name) != nil
}
