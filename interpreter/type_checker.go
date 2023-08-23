package interpreter

import (
	"github.com/charmbracelet/log"

	protocol "github.com/tliron/glsp/protocol_3_16"

	"arc/ast"
	"arc/utilities"
)

type TypeCheckerInstance struct {
	Scope          *TypeScope
	Env            *Environment
	isTypeChecking bool
}

func NewTypeChecker() *TypeCheckerInstance {
	inst := &TypeCheckerInstance{
		Scope:          NewTypeScope(),
		isTypeChecking: false,
	}
	return inst
}

var TypeChecker = NewTypeChecker()

func (self *TypeCheckerInstance) IsTypeChecking(val bool) {
	self.isTypeChecking = val
	self.Scope.isTypeChecking = val
}

func (self *TypeCheckerInstance) FindType(node ast.Node) ast.Type {
	if !self.isTypeChecking {
		self.Scope.CheckingNode(node)
	}

	var resolveType func(ast.Node) ast.Type
	resolveType = func(node ast.Node) ast.Type {
		switch node := node.(type) {

		case *ast.Literal:
			return node

		case *ast.ReturnStatement:
			if node.Value == nil {
				return nil
			}
			return resolveType(node.Value)

		case *ast.CallExpression:
			recvType := self.FindType(node.Receiver)
			if recvType == nil {
				return nil
			}
			obj, ok := recvType.(*ast.ObjectDeclaration)
			if !ok {
				return nil
			}
			m, ok := obj.Methods[node.Function.Name]
			if !ok {
				return nil
			}
			return self.Env.LookupObject(m.ReturnType.Type)

		// case *ast.Identifier:
		// 	return self.Env.LookupObject(node.Name)

		case *ast.TypedIdentifier:
			bt := node.TypeReference.GetBasicType()
			if bt == nil {
				if node.TypeReference == nil {
					return nil
				}
				return resolveType(node.TypeReference)
			}
			return bt

		case *ast.VarReference:
			v := self.Scope.Lookup(node.Name)
			if v == nil {
				return nil
			}
			return v

		case *ast.TypeReference:
			obj := self.Env.LookupType(node.Type)
			if obj == nil {
				return nil
			}
			return obj

		case *ast.IndexAccessExpression:
			resolved := resolveType(node.Instance)
			if resolved == nil {
				return nil
			}
			return ast.BasicTypes["any"]

		case *ast.FieldAccessExpression:
			resolved := resolveType(node.StructInstance)
			if resolved == nil {
				return nil
			}

			if rt, ok := resolved.(*ast.TypeReference); ok {
				resolved = resolveType(rt)
			}

			var resolvedType ast.Type
			switch resolved := resolved.(type) {
			case *ast.ObjectDeclaration:
				for _, field := range resolved.Fields {
					if field.Name == node.FieldName {
						resolvedType = field.TypeReference
						break
					}
				}
			case *ast.EnumDeclaration:
				for _, field := range resolved.Values {
					if field.Name.Name == node.FieldName {
						resolvedType = field.Type
						break
					}
				}

			default:
				return resolved
			}

			// f := resolveType(resolved)

			return resolvedType

		}

		return nil
	}

	resolved := resolveType(node)

	if self.isTypeChecking {
		self.Scope.LinkNodeScope(node)
		if resolved != nil {
			self.Scope.LinkNodeScope(resolved)
		}
	}

	return resolved
}

func (self *TypeCheckerInstance) FindDeclaration(node ast.Node) (ast.Type, ast.Node) {
	if !self.isTypeChecking {
		self.Scope.CheckingNode(node)
	}

	var resolveDeclaration func(ast.Node) (ast.Type, ast.Node)
	resolveDeclaration = func(node ast.Node) (ast.Type, ast.Node) {
		switch node := node.(type) {

		case *ast.Literal:
			return node, node
		case *ast.ReturnStatement:
			if node.Value == nil {
				return nil, nil
			}
			return resolveDeclaration(node.Value)
		case *ast.CallExpression:
			recvType := self.FindType(node.Receiver)
			if recvType == nil {
				return nil, nil
			}
			switch recvNode := recvType.(type) {
			case ast.Declaration:
				methods := recvNode.GetMethods()
				m, ok := methods[node.Function.Name]
				if ok {
					return m, m
				}
			}
			return nil, nil
		case *ast.TypedIdentifier:
			bt := node.TypeReference.GetBasicType()
			if bt == nil {
				if node.TypeReference == nil {
					return nil, nil
				}
				return self.FindType(node.TypeReference), nil
			}
			return bt, nil

		case *ast.VarReference:
			v := self.Scope.Lookup(node.Name)
			if v == nil {
				return nil, nil
			}
			return v, nil

		case *ast.TypeReference:
			if t := self.Env.LookupType(node.Type); t != nil {
				if obj, ok := t.(*ast.ObjectDeclaration); ok {
					return obj, obj.Name
				} else if e, ok := t.(*ast.EnumDeclaration); ok {
					return e, e.Name
				}
				return t, t
			}
			return nil, nil
			// obj := self.Env.LookupObject(node.Type)
			// if obj == nil {
			// 	return nil, nil
			// }
			// return obj, obj.Name

		case *ast.FieldAccessExpression:
			resolved, _ := resolveDeclaration(node.StructInstance)
			if resolved == nil {
				return nil, nil
			}

			if rt, ok := resolved.(*ast.TypeReference); ok {
				resolved, _ = resolveDeclaration(rt)
			}

			var resolvedNode ast.Node
			switch resolved := resolved.(type) {
			case *ast.ObjectDeclaration:
				for _, field := range resolved.Fields {
					if field.Name == node.FieldName {
						resolvedNode = field
						break
					}
				}
			}

			return resolved, resolvedNode
		}

		return nil, nil
	}

	resolved, resolvedNode := resolveDeclaration(node)

	if self.isTypeChecking {
		self.Scope.LinkNodeScope(node)
		if resolved != nil {
			self.Scope.LinkNodeScope(resolved)
		}
	}

	return resolved, resolvedNode
}

func (self *TypeCheckerInstance) DoEqual(lhs, rhs ast.Node) (bool, ast.Type, ast.Type) {
	lhsType := self.FindType(lhs)
	rhsType := self.FindType(rhs)
	if rhsType == nil || lhsType == nil {
		lhsType = self.FindType(lhs)
		rhsType = self.FindType(rhs)
		return false, nil, nil
	}

	lName := lhsType.TypeName()
	if lName == "any" {
		return true, lhsType, rhsType
	}

	rName := rhsType.TypeName()

	if rName == lName {
		return true, lhsType, rhsType
	}

	return false, lhsType, rhsType
}

func (self *TypeCheckerInstance) MustEqual(mainNode ast.Node, lhs, rhs ast.Node) bool {

	doEqual, lhsType, rhsType := self.DoEqual(lhs, rhs)

	if !doEqual {
		lhsName := ""
		rhsName := ""
		if lhsType != nil {
			lhsName = lhsType.TypeName()
		} else {
			if lhsT, ok := lhs.(ast.Type); ok {
				lhsName = lhsT.TypeName()
			}
			if lhsTok := lhs.GetToken(); lhsTok != nil {
				lhsName = lhsTok.Value
			}
		}

		if rhsType != nil {
			rhsName = rhsType.TypeName()
		} else {
			if rhsT, ok := rhs.(ast.Type); ok {
				rhsName = rhsT.TypeName()
			}
			if rhsTok := rhs.GetToken(); rhsTok != nil {
				rhsName = rhsTok.Value
			}
		}

		NewErrorAtNode(mainNode, "Expected type '%s', but '%s' given", lhsName, rhsName)
	}

	return doEqual

}

// node = the node we'll breadth-first search from(usually ast.Program)
func (self *TypeCheckerInstance) GetNodeAtPosition(node ast.Node, position protocol.Position) ast.Node {
	timer := utilities.NewTimer("GetNodeAtPosition - " + node.GetToken().String())
	defer timer.StopAndLog()

	visited := make(map[ast.Node]bool)

	var visitChildren func(node ast.Node, visitedNodes map[ast.Node]bool) ast.Node
	visitChildren = func(node ast.Node, visitedNodes map[ast.Node]bool) ast.Node {
		if node == nil || !node.IsNodeValid() {
			return nil
		}

		if visitedNodes[node] {
			return nil
		}

		visitedNodes[node] = true

		ctx := node.GetToken()
		var tokenRange *ast.TokenRange
		if ctx != nil {
			tokenRange = node.GetTokenRange()

			startLine := max(0, protocol.UInteger(tokenRange.StartLine)-1)
			startColumn := max(0, protocol.UInteger(tokenRange.RangeStartCol))

			stopLine := max(0, protocol.UInteger(tokenRange.StopLine)-1)
			stopColumn := protocol.UInteger(tokenRange.RangeStopCol)

			// Check if the provided line/column is within the range of this node
			if startLine <= position.Line && stopLine >= position.Line &&
				startColumn <= position.Character && stopColumn >= position.Character {
				return node
			}
		}

		// If not, continue traversing the AST for children nodes
		children := node.GetChildren()
		for _, child := range children {
			if child == nil {
				continue
			}

			foundNode := visitChildren(child, visitedNodes)
			if foundNode != nil {
				return foundNode
			}
		}
		return nil
	}

	return visitChildren(node, visited)
}

// Get the root of the given node and find the corresponding source file
// Used for the LSP textDocument/definition lookup
func (self *TypeCheckerInstance) GetNodeSourceFile(node ast.Node) *SourceFile {
	if node == nil {
		return nil
	}

	root := node.GetRoot()
	if root == nil {
		log.Error("Node %v has no root", node)
		return nil
	}

	for _, file := range Engine.SourceFiles {
		if file.Program == root {
			return file
		}
	}

	return nil
}
