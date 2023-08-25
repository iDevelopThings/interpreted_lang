package interpreter

import (
	"arc/ast"
)

type TypeCheckingVisitor struct {
	*ast.NodeVisitorAdapter
	env *Environment
}

func NewTypeCheckingVisitor(program *ast.Program, env *Environment) *TypeCheckingVisitor {
	v := &TypeCheckingVisitor{
		NodeVisitorAdapter: &ast.NodeVisitorAdapter{},
		env:                env,
	}

	TypeChecker.Env = env

	v.Visit(program)

	return v
}

func (self *TypeCheckingVisitor) Visit(node ast.Node) {
	TypeChecker.Scope.LinkNodeScope(node)

	switch node := node.(type) {
	case *ast.Program:
		// self.VisitProgram(node)
		node.Accept(self)
	case *ast.Block:
		// self.VisitBlock(node)
		node.Accept(self)
	case *ast.EnumDeclaration:
		node.Accept(self)
	case *ast.EnumValue:
		node.Accept(self)
	case *ast.Identifier:
		// self.VisitIdentifier(node)
		node.Accept(self)
	case *ast.TypedIdentifier:
		// self.VisitTypedIdentifier(node)
		node.Accept(self)
	case *ast.ObjectDeclaration:
		// self.VisitObjectDeclaration(node)
		node.Accept(self)
	case *ast.FunctionDeclaration:
		// self.VisitFunctionDeclaration(node)
		node.Accept(self)
	case *ast.RangeExpression:
		// self.VisitRangeExpression(node)
		node.Accept(self)
	case *ast.BinaryExpression:
		// self.VisitBinaryExpression(node)
		node.Accept(self)
	case *ast.UnaryExpression:
		// self.VisitUnaryExpression(node)
		node.Accept(self)
	case *ast.FieldAccessExpression:
		// self.VisitFieldAccessExpression(node)
		node.Accept(self)
	case *ast.IndexAccessExpression:
		// self.VisitIndexAccessExpression(node)
		node.Accept(self)
	case *ast.CallExpression:
		// self.VisitCallExpression(node)
		node.Accept(self)
	case *ast.HttpRouteDeclaration:
		// self.VisitHttpRouteDeclaration(node)
		node.Accept(self)
	case *ast.HttpResponseData:
		// self.VisitHttpResponseData(node)
		node.Accept(self)
	case *ast.HttpRouteBodyInjectionStatement:
		// self.VisitHttpRouteBodyInjection(node)
		node.Accept(self)
	case *ast.Literal:
		// self.VisitLiteral(node)
		node.Accept(self)
	case *ast.ArrayInstantiation:
		// self.VisitArrayInstantiation(node)
		node.Accept(self)
	case *ast.ObjectInstantiation:
		// self.VisitObjectInstantiation(node)
		node.Accept(self)
	case *ast.DictionaryInstantiation:
		// self.VisitDictionaryInstantiation(node)
		node.Accept(self)
	case *ast.IfStatement:
		// self.VisitIfStatement(node)
		node.Accept(self)
	case *ast.LoopStatement:
		// self.VisitLoopStatement(node)
		node.Accept(self)
	case *ast.AssignmentStatement:
		// self.VisitAssignmentStatement(node)
		node.Accept(self)
	case *ast.VarReference:
		// self.VisitVarReference(node)
		node.Accept(self)
	case *ast.ReturnStatement:
		// self.VisitReturnStatement(node)
		node.Accept(self)
	case *ast.BreakStatement:
		// self.VisitBreakStatement(node)
		node.Accept(self)
	case *ast.DeleteStatement:
		// self.VisitDeleteStatement(node)
		node.Accept(self)
	}
}

func (self *TypeCheckingVisitor) VisitProgram(node *ast.Program) {
	TypeChecker.Scope.Push()
}

func (self *TypeCheckingVisitor) VisitEnumDeclaration(node *ast.EnumDeclaration) {
	if TypeChecker.Scope.IsDefined(node.Name.Name) {
		NewErrorAtNode(node.Name, "Enum '%s' already defined", node.Name.Name)
	}
	TypeChecker.Scope.Insert(node.Name, node)
}

func (self *TypeCheckingVisitor) VisitFunctionDeclaration(node *ast.FunctionDeclaration) {
	TypeChecker.Scope.Push()
	defer TypeChecker.Scope.Pop()

	if node.Args != nil && len(node.Args) > 0 {
		for _, arg := range node.Args {
			if self.env.LookupObject(arg.TypeReference.Type) == nil {
				NewErrorAtNode(arg.TypeReference, "Type '%s' is not defined", arg.TypeReference.Type)
			}

			if !TypeChecker.Scope.IsDefined(arg.Name) {
				TypeChecker.Scope.Insert(arg.Identifier, arg.TypeReference.GetBasicType())
			}
		}
	}

	if node.Receiver != nil {
		if node.Receiver.TypeReference != nil {
			if self.env.LookupObject(node.Receiver.TypeReference.Type) == nil {
				NewErrorAtNode(node.Receiver.TypeReference, "Type '%s' is not defined", node.Receiver.TypeReference.Type)
			}
			if !node.IsStatic {
				if !TypeChecker.Scope.IsDefined(node.Receiver.Name) {
					TypeChecker.Scope.Insert(node.Receiver.Identifier, node.Receiver.TypeReference)
				}
			}
		}
	}

	if node.ReturnType == nil {
		NewErrorAtNode(node.ReturnType, "Return type: type '%s' is not defined", node.ReturnType.Type)
	} else {
		if node.ReturnType.Type != "void" {
			if self.env.LookupObject(node.ReturnType.Type) == nil {
				NewErrorAtNode(node.ReturnType, "Return type: type '%s' is not defined", node.ReturnType.Type)
			}
		}
	}

	if node.Body != nil {
		node.Body.Accept(self)
	}

}

func (self *TypeCheckingVisitor) VisitBlock(node *ast.Block) {
	TypeChecker.Scope.LinkNodeScope(node)

	if node.Statements != nil {
		for _, statement := range node.Statements {
			statement.Accept(self)
		}
	}
}

func (self *TypeCheckingVisitor) VisitCallExpression(node *ast.CallExpression) {
	if node == nil {
		return
	}
	TypeChecker.Scope.LinkNodeScope(node)

	if len(node.Args) > 0 {
		for _, arg := range node.Args {
			arg.Accept(self)
		}
	}

	var receiverType ast.Type
	if node.Receiver != nil {
		receiverType, _ = TypeChecker.FindDeclaration(node.Receiver)
	}

	if receiverType != nil {
		if enum, ok := receiverType.(*ast.EnumDeclaration); ok {
			valCtor := enum.GetValueConstructor(node.Function.Name)
			if valCtor == nil {
				NewErrorAtNode(node.Function, "Enum '%s' has no value constructor '%s'", enum.Name.Name, node.Function.Name)
			}

			if len(node.Args) != len(valCtor.Properties) {
				NewErrorAtNode(node.Function, "Enum value constructor '%s' expects %d arguments, but %d given", node.Function.Name, len(valCtor.Properties), len(node.Args))
			}

			return
		}
	}

	lookupName := node.Function.Name
	if receiverType != nil {
		// lookupName = receiverType.(*ast.ObjectDeclaration).Name.Name + "_" + node.Function.Name
		lookupName = receiverType.TypeName() + "_" + node.Function.Name
	}

	fnDecl := self.env.LookupFunction(lookupName)
	if fnDecl == nil {
		NewErrorAtNode(node.Function, "Function '%s' is not defined", node.Function.Name)
	}

	if fnDecl.Args != nil && len(fnDecl.Args) > 0 {

		if !fnDecl.HasVariadicArgs {
			if node.Args == nil || len(node.Args) != len(fnDecl.Args) {
				NewErrorAtNode(
					node,
					"Function '%s' expects %d arguments, but %d given",
					node.Function.Name,
					len(fnDecl.Args),
					len(node.Args),
				)
			}

		}

		var declArg *ast.TypedIdentifier
		for i, arg := range node.Args {
			if declArg == nil || !declArg.TypeReference.IsVariadic {
				if i < len(fnDecl.Args) {
					declArg = fnDecl.Args[i]
				}
			}

			if declArg == nil {
				NewErrorAtNode(arg, "Function declaration '%s' does not have argument at index %d", fnDecl.Name, i)
			}

			TypeChecker.MustEqual(arg, declArg, arg)
		}
	}

}

func (self *TypeCheckingVisitor) VisitIdentifier(node *ast.Identifier) {
	// log.Debugf("TypeCheckingVisitor.VisitIdentifier: %#v", node.GetToken())
}

func (self *TypeCheckingVisitor) VisitVarReference(node *ast.VarReference) {
	TypeChecker.Scope.LinkNodeScope(node)

	// log.Debugf("TypeCheckingVisitor.VisitVarReference: %#v", node.GetToken())

	if !TypeChecker.Scope.IsDefined(node.Name) {
		if node.Name != "fmt" {
			NewErrorAtNode(node, "Variable '%s' is not defined", node.Name)
		}
	}
}

func (self *TypeCheckingVisitor) VisitTypeReference(node *ast.TypeReference) {
	// log.Debugf("TypeCheckingVisitor.VisitTypeReference: %#v", node.GetToken())

	TypeChecker.Scope.LinkNodeScope(node)
}

func (self *TypeCheckingVisitor) VisitAssignmentStatement(node *ast.AssignmentStatement) {
	// log.Debugf("TypeCheckingVisitor.VisitAssignmentStatement: %#v", node.GetToken())

	if node.Name != nil {
		node.Name.Accept(self)
	}

	var varType ast.Type

	if node.Type != nil {
		node.Type.Accept(self)

		varType = self.env.LookupType(node.Type.Type)
		if varType == nil {
			NewErrorAtNode(node.Type, "Type '%s' is not defined", node.Type.Type)
		}
	}

	switch v := node.Value.(type) {
	case *ast.Literal:
		if !TypeChecker.Scope.IsDefined(node.Name.Name) {
			bt := node.Type.GetBasicType()
			if bt == nil {
				NewErrorAtNode(node, "[VisitAssignmentStatement-Literal]: Type '%s' is not defined", node.Type.Type)
			}
			TypeChecker.Scope.Insert(node.Name, bt)
		}

	case *ast.VarReference:
		v.Accept(self)

	case *ast.ObjectInstantiation:
		objDecl := self.env.LookupObject(v.TypeName.Name)
		if objDecl == nil {
			NewErrorAtNode(v, "[VisitAssignmentStatement-ObjectInstantiation]: Type '%s' is not defined", v.TypeName.Name)
		}

		if node.Type.Type != objDecl.TypeName() {
			NewErrorAtNode(node, "[VisitAssignmentStatement-ObjectInstantiation]: Type '%s' does not match defined type of '%s'", node.Type.Type, objDecl.TypeName())
		}

		if !TypeChecker.Scope.IsDefined(node.Name.Name) {
			TypeChecker.Scope.Insert(node.Name, objDecl)
		}

	case *ast.DictionaryInstantiation:
		if node.Type.Type == "" {
			node.Type.Type = "dict"
		}

		if !TypeChecker.Scope.IsDefined(node.Name.Name) {
			TypeChecker.Scope.Insert(node.Name, node.Type)
		}

	case *ast.FieldAccessExpression:
		result := TypeChecker.FindType(v)
		if result == nil {
			NewErrorAtNode(v, "[VisitAssignmentStatement-FieldAccessExpression]: Failed to infer type of '%s'", v.GetToken())
		}

		if node.Type.Type == "" {
			node.Type.Type = result.TypeName()
		} else {
			if node.Type.Type != result.TypeName() {
				NewErrorAtNode(node, "[VisitAssignmentStatement-FieldAccessExpression]: Type '%s' does not match defined type of '%s'", node.Type.Type, result.TypeName())
			}
		}

		if !TypeChecker.Scope.IsDefined(node.Name.Name) {
			TypeChecker.Scope.Insert(node.Name, result)
		}

	case *ast.CallExpression:
		result := TypeChecker.FindType(v.Receiver)
		switch result.(type) {
		// Our enum value `constructor`s are seen as call expressions,
		// so we need to handle them here
		case *ast.EnumDeclaration:
			if !TypeChecker.Scope.IsDefined(node.Name.Name) {
				TypeChecker.Scope.Insert(node.Name, result)
			}
		default:
			if result == nil {
				NewErrorAtNode(v, "[VisitAssignmentStatement-CallExpression]: Failed to infer type of '%s'", v.GetToken())
			}

			NewErrorAtNode(node, "[VisitAssignmentStatement-CallExpression]: I don't know what dragons lay here... %T", result)
		}

	default:
		if node.Type.AstNode != nil {
			NewErrorAtNode(node.Type, "[VisitAssignmentStatement-default]:Type '%s' is not defined: %#v", node.Type.Type, v)
			return
		}
		NewErrorAtNode(node, "[VisitAssignmentStatement-default]:Type '%s' is not defined: %#v", node.Type.Type, v)
	}

}

func (self *TypeCheckingVisitor) VisitTypedIdentifier(node *ast.TypedIdentifier) {
	TypeChecker.Scope.LinkNodeScope(node)

	// log.Debugf("TypeCheckingVisitor.VisitTypedIdentifier: %#v", node.GetToken())

	if node.TypeReference != nil {
		node.TypeReference.Accept(self)
	}
}

func (self *TypeCheckingVisitor) VisitReturnStatement(node *ast.ReturnStatement) {
	// log.Debugf("TypeCheckingVisitor.VisitReturnStatement: %#v", node)

	TypeChecker.Scope.LinkNodeScope(node)

	funcDecl := ast.FindFirstParentOfType[*ast.FunctionDeclaration](node)

	if funcDecl == nil {
		NewErrorAtNode(node, "Return statement must be inside a function")
	}

	if funcDecl.ReturnType == nil {
		NewErrorAtNode(node, "Failed to get return type of function declaration, this is likely a compiler bug")
	}

	if funcDecl.ReturnType.Type == "void" {
		if node.Value != nil {
			NewErrorAtNode(node, "Cannot return a value from a function with void return type")
		}
	} else {
		if node.Value == nil {
			NewErrorAtNode(node, "Must return a value from a function with non-void return type")
		}
	}

	// TODO: Compare the type of the return statement value to the function return type

	switch v := node.Value.(type) {
	case *ast.VarReference:
		if !TypeChecker.Scope.IsDefined(v.Name) {
			NewErrorAtNode(v, "[VisitReturnStatement-VarReference]: Type '%s' is not defined", v.Name)
		}

	case *ast.ObjectInstantiation:
		objDecl := self.env.LookupObject(v.TypeName.Name)
		if objDecl == nil {
			NewErrorAtNode(v, "[VisitReturnStatement-ObjectInstantiation]: Type '%s' is not defined", v.TypeName.Name)
		}

	case *ast.FieldAccessExpression:
		result := TypeChecker.FindType(v)
		if result == nil {
			NewErrorAtNode(v, "[VisitReturnStatement-FieldAccessExpression]: Failed to infer type of '%s'", v.GetToken())
		}
	case *ast.Literal:
		valType := v.GetBasicType()
		switch {

		case valType == ast.NoneType:
			if !funcDecl.ReturnType.IsOptionType {
				newRt := "?" + funcDecl.ReturnType.Type
				ErrorManager.ShouldExit(false)

				NewErrorAtNode(node.Value, "You cannot return 'none', without your value being an option type.")
				// NewErrorAtNode(funcDecl.ReturnType, "You cannot compare a non-option type to 'none', you should change your function return type to: %s", newRt)
				ErrorManager.ShouldExit(true)
				NewErrorAtNode(funcDecl.ReturnType, "You should change your return type to: %s", newRt)
			}

		default:
			if funcDecl.ReturnType.Type != v.TypeName() {
				NewErrorAtNode(node, "[VisitReturnStatement-Literal]: Type '%s' does not match defined type of '%s'", v.TypeName(), funcDecl.ReturnType.Type)
			}

		}

	default:
		NewErrorAtNode(node, "[VisitReturnStatement-default]: Failed to resolve type of '%s'", node.GetToken())
	}
}

func (self *TypeCheckingVisitor) VisitIndexAccessExpression(node *ast.IndexAccessExpression) {
	// log.Debugf("TypeCheckingVisitor.VisitIndexAccessExpression: %#v", node)

	TypeChecker.Scope.LinkNodeScope(node)

	if node.Instance != nil {
		node.Instance.Accept(self)
	}
	if node.StartIndex != nil {
		node.StartIndex.Accept(self)
	}
	if node.EndIndex != nil {
		node.EndIndex.Accept(self)
	}

	result := TypeChecker.FindType(node.Instance)
	if result == nil {
		NewErrorAtNode(node, "[VisitIndexAccessExpression]: Failed to infer type of '%s'", node.GetToken())
	}

	switch t := result.(type) {
	case *ast.EnumDeclaration:
		if node.IsSlice {
			NewErrorAtNode(node, "[VisitIndexAccessExpression]: Cannot use slice syntax on enum value access '%s'", t.Name)
		}
		index := node.StartIndex
		if index == nil {
			NewErrorAtNode(node, "[VisitIndexAccessExpression]: Enum value access '%s' requires an index", t.Name)
		}
		indexType := TypeChecker.FindType(index)
		if indexType == nil {
			NewErrorAtNode(node, "[VisitIndexAccessExpression]: Failed to infer type of '%s'", index.GetToken())
		}

		switch indexType.(type) {
		case *ast.BasicType, *ast.Literal:
		default:
			NewErrorAtNode(node, "[VisitIndexAccessExpression]: Enum value access '%s' can only use literal values for item lookup", t.Name)
		}

	}

	// if result.TypeName() != "list" {
	// 	NewErrorAtNode(node, "[VisitIndexAccessExpression]: Type '%s' does not support index access", result.TypeName())
	// }
}
