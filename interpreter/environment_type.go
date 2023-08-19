package interpreter

import (
	"github.com/charmbracelet/log"

	"arc/ast"
)

type TypeScopeMap = map[string]ast.Type

type TypeScope struct {
	Types []TypeScopeMap

	// Node Global ID -> TypeScopeMap
	NodeScope map[int64]TypeScopeMap

	// After type checking has finished, we don't have the `Types` stack anymore.
	// So when we want to check a node, we call `CheckingNode(Node)` to set the current
	// scope to the scope of the node.
	tempCurrentScope TypeScopeMap

	isTypeChecking bool
}

func NewTypeScope() *TypeScope {
	return &TypeScope{
		Types:     []TypeScopeMap{},
		NodeScope: map[int64]TypeScopeMap{},
	}
}

func (self *TypeScope) Push() {
	self.Types = append(self.Types, TypeScopeMap{})
}

func (self *TypeScope) Pop() {
	self.Types = self.Types[:len(self.Types)-1]
}

func (self *TypeScope) Insert(identifier *ast.Identifier, typ ast.Type) {
	self.Types[len(self.Types)-1][identifier.Name] = typ
	self.LinkNodeIdentifierScope(identifier, typ)
}

func (self *TypeScope) LinkNodeScope(node ast.Node) {
	if len(self.Types) == 0 {
		return
	}

	nodeId := node.GetId()
	self.NodeScope[nodeId] = self.Types[len(self.Types)-1]
}
func (self *TypeScope) LinkNodeIdentifierScope(identifier *ast.Identifier, typ ast.Type) {
	self.NodeScope[identifier.GetId()] = self.Types[len(self.Types)-1]
	self.LinkNodeScope(typ)
}

func (self *TypeScope) Lookup(name string) ast.Type {
	if !self.isTypeChecking {
		if self.tempCurrentScope != nil {
			if typ, ok := self.tempCurrentScope[name]; ok {
				return typ
			}
		}
		return nil
	}

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

func (self *TypeScope) CheckingNode(node ast.Node) {
	if node == nil {
		return
	}

	s, ok := self.NodeScope[node.GetId()]
	if !ok {
		self.tempCurrentScope = nil
		log.Warnf("Node %d(value=%v) does not have an assigned node scope", node.GetId(), node)
		return
	}

	self.tempCurrentScope = s
}
