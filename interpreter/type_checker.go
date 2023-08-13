package interpreter

import (
	"interpreted_lang/ast"
)

type TypeCheckerInstance struct {
	Scope *TypeScope
	Env   *Environment
}

func (self *TypeCheckerInstance) FindType(node ast.Node) ast.Type {
	switch node := node.(type) {

	case *ast.Literal:
		return node
	case *ast.TypedIdentifier:
		bt := node.TypeReference.GetBasicType()
		if bt == nil {
			return nil
		}
		return bt

	case *ast.VarReference:
		v := self.Scope.Lookup(node.Name)
		if v == nil {
			return nil
		}
		return v

	case *ast.TypeReference:
		obj := self.Env.LookupObject(node.Type)
		if obj == nil {
			return nil
		}
		return obj

	case *ast.FieldAccessExpression:
		resolved := self.FindType(node.StructInstance)
		if resolved == nil {
			return nil
		}

		if rt, ok := resolved.(*ast.TypeReference); ok {
			resolved = self.FindType(rt)
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

		default:
			return resolved
		}

		// f := self.FindType(resolved)

		return resolvedType

	}

	return nil
}

func (self *TypeCheckerInstance) DoEqual(lhs, rhs ast.Node) (bool, ast.Type, ast.Type) {
	lhsType := self.FindType(lhs)
	rhsType := self.FindType(rhs)
	if rhsType == nil || lhsType == nil {
		return false, nil, nil
	}

	if lhsType.TypeName() == rhsType.TypeName() {
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
			lhsName = lhs.GetRule().GetText()
		}

		if rhsType != nil {
			rhsName = rhsType.TypeName()
		} else {
			rhsName = rhs.GetRule().GetText()
		}

		NewErrorAtNode(mainNode, "Expected type '%s', but '%s' given", lhsName, rhsName)
	}

	return doEqual

}

func NewTypeChecker() *TypeCheckerInstance {
	inst := &TypeCheckerInstance{
		Scope: NewTypeScope(),
	}
	return inst
}

var TypeChecker = NewTypeChecker()
