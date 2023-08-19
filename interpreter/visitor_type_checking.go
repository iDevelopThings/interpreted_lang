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
	case *ast.AssignmentExpression:
		// self.VisitAssignmentExpression(node)
		node.Accept(self)
	case *ast.BinaryExpression:
		// self.VisitBinaryExpression(node)
		node.Accept(self)
	case *ast.PostfixExpression:
		// self.VisitPostfixExpression(node)
		node.Accept(self)
	case *ast.UnaryExpression:
		// self.VisitUnaryExpression(node)
		node.Accept(self)
	case *ast.FieldAccessExpression:
		// self.VisitFieldAccessExpression(node)
		node.Accept(self)
	case *ast.IndexAccessExpression:
		// self.VisitArrayAccessExpression(node)
		node.Accept(self)
	case *ast.CallExpression:
		// self.VisitCallExpression(node)
		node.Accept(self)
	case *ast.HttpRouteDeclaration:
		// self.VisitHttpRouteDeclaration(node)
		node.Accept(self)
	case *ast.HttpServerConfig:
		// self.VisitHttpServerConfig(node)
		node.Accept(self)
	case *ast.HttpResponseData:
		// self.VisitHttpResponseData(node)
		node.Accept(self)
	case *ast.HttpRouteBodyInjection:
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
		NewErrorAtNode(node.ReturnType, "Return type: type '%s' is not defined", node.ReturnType.Name)
	} else {
		if node.ReturnType.Name != "void" {
			if self.env.LookupObject(node.ReturnType.Name) == nil {
				NewErrorAtNode(node.ReturnType, "Return type: type '%s' is not defined", node.ReturnType.Name)
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

	var receiverType ast.Type
	if node.Receiver != nil {
		receiverType, _ = TypeChecker.FindDeclaration(node.Receiver)
	}

	lookupName := node.Function.Name
	if receiverType != nil {
		lookupName = receiverType.(*ast.ObjectDeclaration).Name.Name + "_" + node.Function.Name
	}

	fnDecl := self.env.LookupFunction(lookupName)
	if fnDecl == nil {
		NewErrorAtNode(node, "Function '%s' is not defined", node.Function.Name)
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
	// log.Debugf("TypeCheckingVisitor.VisitIdentifier: %v", node.GetToken())
}

func (self *TypeCheckingVisitor) VisitVarReference(node *ast.VarReference) {
	TypeChecker.Scope.LinkNodeScope(node)

	// log.Debugf("TypeCheckingVisitor.VisitVarReference: %v", node.GetToken())

	if !TypeChecker.Scope.IsDefined(node.Name) {
		if node.Name != "fmt" {
			NewErrorAtNode(node, "Variable '%s' is not defined", node.Name)
		}
	}
}

func (self *TypeCheckingVisitor) VisitTypeReference(node *ast.TypeReference) {
	// log.Debugf("TypeCheckingVisitor.VisitTypeReference: %v", node.GetToken())

	TypeChecker.Scope.LinkNodeScope(node)
}

func (self *TypeCheckingVisitor) VisitAssignmentStatement(node *ast.AssignmentStatement) {
	// log.Debugf("TypeCheckingVisitor.VisitAssignmentStatement: %v", node.GetToken())

	if node.Name != nil {
		node.Name.Accept(self)
	}
	if node.Type != nil {
		node.Type.Accept(self)
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

	default:
		NewErrorAtNode(node, "[VisitAssignmentStatement-default]:Type '%s' is not defined: %v", node.Type.Type, v)
	}

}

func (self *TypeCheckingVisitor) VisitTypedIdentifier(node *ast.TypedIdentifier) {
	TypeChecker.Scope.LinkNodeScope(node)

	// log.Debugf("TypeCheckingVisitor.VisitTypedIdentifier: %v", node.GetToken())

	if node.TypeReference != nil {
		node.TypeReference.Accept(self)
	}
}

func (self *TypeCheckingVisitor) VisitReturnStatement(node *ast.ReturnStatement) {
	// log.Debugf("TypeCheckingVisitor.VisitReturnStatement: %v", node)

	TypeChecker.Scope.LinkNodeScope(node)

	funcDecl := ast.FindFirstParentOfType[*ast.FunctionDeclaration](node)

	if funcDecl == nil {
		NewErrorAtNode(node, "Return statement must be inside a function")
	}

	if funcDecl.ReturnType == nil {
		NewErrorAtNode(node, "Failed to get return type of function declaration, this is likely a compiler bug")
	}

	if funcDecl.ReturnType.Name == "void" {
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
		if funcDecl.ReturnType.Name != v.TypeName() {
			NewErrorAtNode(node, "[VisitReturnStatement-Literal]: Type '%s' does not match defined type of '%s'", v.TypeName(), funcDecl.ReturnType.Name)
		}

	default:
		NewErrorAtNode(node, "[VisitReturnStatement-default]: Failed to resolve type of '%s'", node.GetToken())
	}
}
