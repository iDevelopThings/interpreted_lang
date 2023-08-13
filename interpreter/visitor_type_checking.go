package interpreter

import (
	"github.com/charmbracelet/log"

	"interpreted_lang/ast"
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
	case *ast.ArrayAccessExpression:
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
				TypeChecker.Scope.Insert(arg.Name, arg.TypeReference.GetBasicType())
			}
		}
	}

	if node.Receiver != nil {
		if self.env.LookupObject(node.Receiver.TypeReference.Type) == nil {
			NewErrorAtNode(node.Receiver.TypeReference, "Type '%s' is not defined", node.Receiver.TypeReference.Type)
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
	log.Debugf("TypeCheckingVisitor.VisitBlock: %v", node)
}

func (self *TypeCheckingVisitor) VisitCallExpression(node *ast.CallExpression) {
	if node == nil {
		return
	}

	fnDecl := self.env.LookupFunction(node.FunctionName)
	if fnDecl == nil {
		NewErrorAtNode(node, "Function '%s' is not defined", node.FunctionName)
	}

	if fnDecl.Args != nil && len(fnDecl.Args) > 0 {
		if node.Args == nil || len(node.Args) != len(fnDecl.Args) {
			NewErrorAtToken(
				node.ArgsToken,
				"Function '%s' expects %d arguments, but %d given",
				node.FunctionName,
				len(fnDecl.Args),
				len(node.Args),
			)
		}

		for i, fnCallArg := range node.Args {
			declArg := fnDecl.Args[i]

			TypeChecker.MustEqual(fnCallArg, declArg, fnCallArg)

		}
	}

}

func (self *TypeCheckingVisitor) VisitIdentifier(node *ast.Identifier) {
	// log.Debugf("TypeCheckingVisitor.VisitIdentifier: %v", node.GetRule().GetText())
}

func (self *TypeCheckingVisitor) VisitVarReference(node *ast.VarReference) {
	// log.Debugf("TypeCheckingVisitor.VisitVarReference: %v", node.GetRule().GetText())

	if !TypeChecker.Scope.IsDefined(node.Name) {
		NewErrorAtNode(node, "Variable '%s' is not defined", node.Name)
	}
}

func (self *TypeCheckingVisitor) VisitAssignmentStatement(node *ast.AssignmentStatement) {
	// log.Debugf("TypeCheckingVisitor.VisitAssignmentStatement: %v", node.GetRule().GetText())

	switch v := node.Value.(type) {
	case *ast.Literal:
		if !TypeChecker.Scope.IsDefined(node.TypedIdentifier.Name) {
			bt := node.TypedIdentifier.TypeReference.GetBasicType()
			if bt == nil {
				NewErrorAtNode(node, "[VisitAssignmentStatement-Literal]: Type '%s' is not defined", node.TypedIdentifier.TypeReference.Type)
			}
			TypeChecker.Scope.Insert(node.TypedIdentifier.Name, bt)
		}

	case *ast.VarReference:
		if !TypeChecker.Scope.IsDefined(v.Name) {
			NewErrorAtNode(v, "[VisitAssignmentStatement-VarReference]: Type '%s' is not defined", v.Name)
		}

	case *ast.ObjectInstantiation:
		objDecl := self.env.LookupObject(v.TypeName)
		if objDecl == nil {
			NewErrorAtNode(v, "[VisitAssignmentStatement-ObjectInstantiation]: Type '%s' is not defined", v.TypeName)
		}

		if node.TypedIdentifier.TypeReference.Type != objDecl.TypeName() {
			NewErrorAtNode(node, "[VisitAssignmentStatement-ObjectInstantiation]: Type '%s' does not match defined type of '%s'", node.TypedIdentifier.TypeReference.Type, objDecl.TypeName())
		}

		if !TypeChecker.Scope.IsDefined(node.TypedIdentifier.Name) {
			TypeChecker.Scope.Insert(node.TypedIdentifier.Name, objDecl)
		}

	case *ast.DictionaryInstantiation:
		if node.TypedIdentifier.TypeReference.Type == "" {
			node.TypedIdentifier.TypeReference.Type = "dict"
		}

		if !TypeChecker.Scope.IsDefined(node.TypedIdentifier.Name) {
			TypeChecker.Scope.Insert(node.TypedIdentifier.Name, node.TypeReference)
		}

	case *ast.FieldAccessExpression:
		result := TypeChecker.FindType(v)
		if result == nil {
			NewErrorAtNode(v, "[VisitAssignmentStatement-FieldAccessExpression]: Failed to infer type of '%s'", v.GetRule().GetText())
		}

		if node.TypedIdentifier.TypeReference.Type == "" {
			node.TypedIdentifier.TypeReference.Type = result.TypeName()
		} else {
			if node.TypedIdentifier.TypeReference.Type != result.TypeName() {
				NewErrorAtNode(node, "[VisitAssignmentStatement-FieldAccessExpression]: Type '%s' does not match defined type of '%s'", node.TypedIdentifier.TypeReference.Type, result.TypeName())
			}
		}

		if !TypeChecker.Scope.IsDefined(node.TypedIdentifier.Name) {
			TypeChecker.Scope.Insert(node.TypedIdentifier.Name, result)
		}

	default:
		NewErrorAtNode(node, "[VisitAssignmentStatement-default]:Type '%s' is not defined: %v", node.TypedIdentifier.TypeReference.Type, v)
	}

}

func (self *TypeCheckingVisitor) VisitTypedIdentifier(node *ast.TypedIdentifier) {
}
