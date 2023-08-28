package interpreter

import (
	"github.com/charmbracelet/log"

	"arc/ast"
)

var Inference *InferenceInstance = NewInference()

func NewInference() *InferenceInstance {
	inst := &InferenceInstance{}

	return inst
}

type InferenceInstance struct {
}

// typeCheckerLookupCb is purely used for the TypeChecker
// With some nodes like VarReferences, we need to look it up in the type checker scope
// Which we can't do properly here...
func (self *InferenceInstance) FindType(node ast.Node, typeCheckerLookupCb ...func(ast.Node) ast.Type) ast.Type {
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
			return Registry.LookupObject(m.ReturnType.Type)

		// case *ast.Identifier:
		// 	return Registry.LookupObject(node.Name)

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
			if len(typeCheckerLookupCb) > 0 {
				return typeCheckerLookupCb[0](node)
			} else {
				log.Warnf("[inference] Cannot lookup/resolve, type checker callback is not provided")
				return nil
			}

		case *ast.TypeReference:
			obj := Registry.LookupType(node.Type)
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

			return resolvedType
		}

		return nil
	}

	resolved := resolveType(node)

	return resolved
}

func (self *InferenceInstance) FindDeclaration(node ast.Node, typeCheckerLookupCb ...func(ast.Node) (ast.Type, ast.Node)) (ast.Type, ast.Node) {

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
			if len(typeCheckerLookupCb) > 0 {
				return typeCheckerLookupCb[0](node)
			} else {
				log.Warnf("[inference] Cannot lookup/resolve, type checker callback is not provided")
				return nil, nil
			}

		case *ast.TypeReference:
			if t := Registry.LookupType(node.Type); t != nil {
				if obj, ok := t.(*ast.ObjectDeclaration); ok {
					return obj, obj.Name
				} else if e, ok := t.(*ast.EnumDeclaration); ok {
					return e, e.Name
				}
				return t, t
			}
			return nil, nil

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

	return resolved, resolvedNode
}

func (self *InferenceInstance) GetCallExpressionFunctionDeclaration(node *ast.CallExpression) *ast.FunctionDeclaration {

	var receiverType ast.Type
	if node.Receiver != nil {
		receiverType, _ = TypeChecker.FindDeclaration(node.Receiver)
	}

	lookupName := node.Function.Name
	if receiverType != nil {
		lookupName = receiverType.TypeName() + "_" + node.Function.Name
	}

	return Registry.LookupFunction(lookupName)
}

func (self *InferenceInstance) InferExpressionType(n ast.Node) ast.Type {

	switch node := n.(type) {

	case *ast.CallExpression:
		{
			fnDecl := self.GetCallExpressionFunctionDeclaration(node)
			if fnDecl == nil {
				return nil
			}

			return fnDecl.ReturnType
		}

	case *ast.Literal:
		return node.GetBasicType()

	case *ast.VarReference:
		t, _ := self.FindDeclaration(node)
		return t

	default:
		log.Warnf("[inference] Unhandled node type: %T", node)
		return nil
	}
}
