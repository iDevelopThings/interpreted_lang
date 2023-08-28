package interpreter

import (
	"arc/ast"
	"arc/interpreter/diagnostics"
	"arc/log"
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

	resolved := Inference.FindType(node, func(n ast.Node) ast.Type {
		switch node := n.(type) {

		case *ast.VarReference:
			v := self.Scope.Lookup(node.Name)
			if v == nil {
				return nil
			}
			return v

		default:
			log.Warnf("TypeChecker.FindType: unhandled node type %T", node)
			return nil
		}
	})

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

	resolved, resolvedNode := Inference.FindDeclaration(node, func(n ast.Node) (ast.Type, ast.Node) {
		switch node := n.(type) {
		case *ast.VarReference:
			v := self.Scope.Lookup(node.Name)
			if v == nil {
				return nil, nil
			}
			return v, nil
		default:
			log.Warnf("TypeChecker.FindDeclaration: unhandled node type %T", node)
			return nil, nil
		}
	})

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

func (self *TypeCheckerInstance) TypeCheckTree(program *ast.Program, env *Environment) {
	self.Env = env

	ast.WalkWithVisitEvent(program, func(node ast.Node) (bool, ast.VisitFuncLeaveCallback) {

		switch node := node.(type) {
		case *ast.Program:
			self.typeCheckProgram(node)
			return true, nil
		case *ast.EnumDeclaration:
			self.typeCheckEnumDeclaration(node)
			return true, nil
		case *ast.FunctionDeclaration:
			leaveCb := self.typeCheckFunctionDeclaration(node)
			return true, leaveCb
		case *ast.Block:
			self.typeCheckBlock(node)
			return true, nil
		case *ast.CallExpression:
			self.typeCheckCallExpression(node)
			return true, nil
		case *ast.Identifier:
			self.typeCheckIdentifier(node)
			return true, nil
		case *ast.VarReference:
			self.typeCheckVarReference(node)
			return true, nil
		case *ast.FieldAccessExpression:
			self.typeCheckFieldAccessExpression(node)
			return true, nil
		case *ast.TypeReference:
			self.typeCheckTypeReference(node)
			return true, nil
		case *ast.BinaryExpression:
			self.typeCheckBinaryExpression(node)
			return true, nil
		case *ast.AssignmentStatement:
			self.typeCheckAssignmentStatement(node)
			return true, nil
		case *ast.TypedIdentifier:
			self.typeCheckTypedIdentifier(node)
			return true, nil
		case *ast.ReturnStatement:
			self.typeCheckReturnStatement(node)
			return true, nil
		case *ast.IndexAccessExpression:
			self.typeCheckIndexAccessExpression(node)
			return true, nil
		}

		return true, nil
	})

}

func (self *TypeCheckerInstance) typeCheckProgram(node *ast.Program) {
	self.Scope.Push()
}
func (self *TypeCheckerInstance) typeCheckEnumDeclaration(node *ast.EnumDeclaration) {
	if self.Scope.IsDefined(node.Name.Name) {
		NewErrorAtNode(node.Name, "Enum '%s' already defined", node.Name.Name)
	}
	self.Scope.Insert(node.Name, node)
}
func (self *TypeCheckerInstance) typeCheckFunctionDeclaration(node *ast.FunctionDeclaration) ast.VisitFuncLeaveCallback {
	self.Scope.Push()

	if node.Args != nil && len(node.Args) > 0 {
		for _, arg := range node.Args {
			if Registry.LookupObject(arg.TypeReference.Type) == nil {
				NewErrorAtNode(arg.TypeReference, "Type '%s' is not defined", arg.TypeReference.Type)
			}

			if !self.Scope.IsDefined(arg.Name) {
				self.Scope.Insert(arg.Identifier, arg.TypeReference.GetBasicType())
			}
		}
	}

	if node.Receiver != nil {
		if node.Receiver.TypeReference != nil {
			if Registry.LookupObject(node.Receiver.TypeReference.Type) == nil {
				NewErrorAtNode(node.Receiver.TypeReference, "Type '%s' is not defined", node.Receiver.TypeReference.Type)
			}
			if !node.IsStatic {
				if !self.Scope.IsDefined(node.Receiver.Name) {
					self.Scope.Insert(node.Receiver.Identifier, node.Receiver.TypeReference)
				}
			}
		}
	}

	if node.ReturnType == nil {
		NewErrorAtNode(node.ReturnType, "Return type: type '%s' is not defined", node.ReturnType.Type)
	} else {
		if node.ReturnType.Type != "void" {
			if Registry.LookupObject(node.ReturnType.Type) == nil {
				NewErrorAtNode(node.ReturnType, "Return type: type '%s' is not defined", node.ReturnType.Type)
			}
		}
	}

	return func(node ast.Node) {
		self.Scope.Pop()
	}
}
func (self *TypeCheckerInstance) typeCheckBlock(node *ast.Block) {
	self.Scope.LinkNodeScope(node)
}
func (self *TypeCheckerInstance) typeCheckCallExpression(node *ast.CallExpression) {
	if node == nil {
		return
	}
	self.Scope.LinkNodeScope(node)

	fnDecl := Inference.GetCallExpressionFunctionDeclaration(node)
	if fnDecl == nil {
		NewDiagnosticAtNode(node.Function, diagnostics.FunctionNotDefined, node.Function.Name)
		return
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

			if len(node.ArgumentList.Entries) != len(valCtor.Properties) {
				NewErrorAtNode(node.Function, "Enum value constructor '%s' expects %d arguments, but %d given", node.Function.Name, len(valCtor.Properties), len(node.ArgumentList.Entries))
			}

			return
		}
	}

	// So... we can get the lengths of the node args & decl args
	// We'll then create a list with the maximum size of the two
	// We can then go over function delcaration args... and assign any info at the correct index
	// Then the same for the call arguments

	if len(node.ArgumentList.Entries) != len(fnDecl.Args) {
		diagnostic := NewMultiDiagnostic()
		extraDiagnosticMeta := map[string]any{}

		maxLen := max(len(node.ArgumentList.Entries), len(fnDecl.Args))
		argumentInfoList := make([]map[string]any, maxLen)

		for i, arg := range fnDecl.Args {
			argumentInfoList[i] = map[string]any{
				"isDeclared": true,
				"name":       arg.Name,
				"type":       arg.TypeReference.Type,
			}
		}

		for i, arg := range node.ArgumentList.Entries {
			argData := map[string]any{
				"isDeclared": false,
				"name":       "",
				"type":       "",
			}
			if argumentInfoList[i] == nil {
				argumentInfoList[i] = argData
			} else {
				argData = argumentInfoList[i]
			}

			if i < len(fnDecl.Args) {
				argData["name"] = fnDecl.Args[i].Name
				argData["type"] = fnDecl.Args[i].TypeReference.Type
			} else {
				argData["isDeclared"] = false

				argType := Inference.InferExpressionType(arg)
				if argType == nil {
					diagnostic.AddError(arg, "Failed to infer type of argument")
					argData["type"] = nil
				} else {
					argData["name"] = argType.TypeName()
					argData["type"] = argType.TypeName()
				}
			}

			argumentInfoList[i] = argData
		}

		extraDiagnosticMeta["argumentInfo"] = argumentInfoList

		diagnostic.AddDiagnostic(
			node.ArgumentList,
			diagnostics.FunctionCallArgCountMismatch,
			node.Function.Name,
			len(fnDecl.Args),
			len(node.ArgumentList.Entries),
		)
		diagnostic.AttachMeta(extraDiagnosticMeta)
		diagnostic.Push()
	}

	var declArg *ast.TypedIdentifier
	for i, arg := range node.ArgumentList.Entries {
		if declArg == nil || !declArg.TypeReference.IsVariadic {
			if i < len(fnDecl.Args) {
				declArg = fnDecl.Args[i]
			}
		}

		if declArg == nil {
			NewErrorAtNode(arg, "Function declaration '%s' does not have argument at index %d", fnDecl.Name, i)
		}

		if i > len(fnDecl.Args)-1 && !declArg.TypeReference.IsVariadic {
			// At this point if we have more args... and the function declaration arg is not variadic
			// Then we have too many args
			break
		}

		TypeChecker.MustEqual(arg, declArg, arg)
	}

	// if !fnDecl.HasVariadicArgs {
	// 	if node.Args == nil || len(node.Args) != len(fnDecl.Args) {
	// 		NewErrorAtNode(
	// 			node,
	// 			"Function '%s' expects %d arguments, but %d given",
	// 			node.Function.Name,
	// 			len(fnDecl.Args),
	// 			len(node.Args),
	// 		)
	// 	}
	// }

}
func (self *TypeCheckerInstance) typeCheckIdentifier(node *ast.Identifier) {
	// log.Debugf("TypeCheckingVisitor.VisitIdentifier: %#v", node.GetToken())
}
func (self *TypeCheckerInstance) typeCheckVarReference(node *ast.VarReference) {
	self.Scope.LinkNodeScope(node)

	if !self.Scope.IsDefined(node.Name) {
		if node.Name != "fmt" {
			NewErrorAtNode(node, "Variable '%s' is not defined", node.Name)
		}
	}
}
func (self *TypeCheckerInstance) typeCheckTypeReference(node *ast.TypeReference) {
	// log.Debugf("TypeCheckingVisitor.VisitTypeReference: %#v", node.GetToken())

	self.Scope.LinkNodeScope(node)
}
func (self *TypeCheckerInstance) typeCheckBinaryExpression(node *ast.BinaryExpression) {

	if node.Left == nil {
		NewErrorAtNode(node, "Binary expression has no left operand")
	}
	if node.Right == nil {
		NewErrorAtNode(node, "Binary expression has no right operand")
	}

	switch node.Kind {
	case ast.BinaryExpressionKindAssignment:
		{
			// leftType, leftNode := TypeChecker.FindDeclaration(node.Left)
			leftType := TypeChecker.FindType(node.Left)
			rightType := TypeChecker.FindType(node.Right)
			if rightType == nil {
				NewErrorAtNode(node.Right, "Failed to resolve right operand of assignment expression")
			}
			if leftType != nil {
				if leftType.TypeName() != rightType.TypeName() {
					NewErrorAtNode(node.Right, "Cannot assign type '%s' to type '%s'", rightType.TypeName(), leftType.TypeName())
				}
				return
			}

			// We can ignore field access expressions, in this case, it's because the field doesn't exist on the object
			// Which will be verified by the FieldAccessExpression type checker
			if _, ok := node.Left.(*ast.FieldAccessExpression); !ok {
				NewErrorAtNode(node.Left, "Failed to resolve left operand of assignment expression")
			}

		}

	}

}
func (self *TypeCheckerInstance) typeCheckAssignmentStatement(node *ast.AssignmentStatement) {
	// log.Debugf("TypeCheckingVisitor.VisitAssignmentStatement: %#v", node.GetToken())

	var varType ast.Type

	if node.Type != nil {
		varType = Registry.LookupType(node.Type.Type)
		if varType == nil {
			NewErrorAtNode(node.Type, "Type '%s' is not defined", node.Type.Type)
		}
	}

	switch v := node.Value.(type) {
	case *ast.Literal:
		if !self.Scope.IsDefined(node.Name.Name) {
			bt := node.Type.GetBasicType()
			if bt == nil {
				NewErrorAtNode(node, "[VisitAssignmentStatement-Literal]: Type '%s' is not defined", node.Type.Type)
			}
			self.Scope.Insert(node.Name, bt)
		}

	case *ast.ObjectInstantiation:
		objDecl := Registry.LookupObject(v.TypeName.Name)
		if objDecl == nil {
			NewErrorAtNode(v, "[VisitAssignmentStatement-ObjectInstantiation]: Type '%s' is not defined", v.TypeName.Name)
		}

		if node.Type.Type != objDecl.TypeName() {
			NewErrorAtNode(node, "[VisitAssignmentStatement-ObjectInstantiation]: Type '%s' does not match defined type of '%s'", node.Type.Type, objDecl.TypeName())
		}

		if !self.Scope.IsDefined(node.Name.Name) {
			self.Scope.Insert(node.Name, objDecl)
		}

	case *ast.DictionaryInstantiation:
		if node.Type.Type == "" {
			node.Type.Type = "dict"
		}

		if !self.Scope.IsDefined(node.Name.Name) {
			self.Scope.Insert(node.Name, node.Type)
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

		if !self.Scope.IsDefined(node.Name.Name) {
			self.Scope.Insert(node.Name, result)
		}

	case *ast.CallExpression:
		result := TypeChecker.FindType(v.Receiver)
		switch result.(type) {
		// Our enum value `constructor`s are seen as call expressions,
		// so we need to handle them here
		case *ast.EnumDeclaration:
			if !self.Scope.IsDefined(node.Name.Name) {
				self.Scope.Insert(node.Name, result)
			}
		default:
			if result == nil {
				NewErrorAtNode(v, "[VisitAssignmentStatement-CallExpression]: Failed to infer type of '%s'", v.GetToken())
			}

			NewErrorAtNode(node, "[VisitAssignmentStatement-CallExpression]: I don't know what dragons lay here... %T", result)
		}

	default:
		if node.Value != nil && varType == nil {
			if node.Type.AstNode != nil {
				NewErrorAtNode(node.Type, "[VisitAssignmentStatement-default]:Type '%s' is not defined: %#v", node.Type.Type, v)
				return
			}
			NewErrorAtNode(node, "[VisitAssignmentStatement-default]:Type '%s' is not defined: %#v", node.Type.Type, v)
		}
	}

}
func (self *TypeCheckerInstance) typeCheckTypedIdentifier(node *ast.TypedIdentifier) {
	self.Scope.LinkNodeScope(node)

	// log.Debugf("TypeCheckingVisitor.VisitTypedIdentifier: %#v", node.GetToken())
}
func (self *TypeCheckerInstance) typeCheckReturnStatement(node *ast.ReturnStatement) {
	// log.Debugf("TypeCheckingVisitor.VisitReturnStatement: %#v", node)

	self.Scope.LinkNodeScope(node)

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
		if !self.Scope.IsDefined(v.Name) {
			NewErrorAtNode(v, "[VisitReturnStatement-VarReference]: Type '%s' is not defined", v.Name)
		}

	case *ast.ObjectInstantiation:
		objDecl := Registry.LookupObject(v.TypeName.Name)
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

				NewMultiDiagnostic().
					AddError(node.Value, "You cannot return 'none' without your value being an option type").
					AddError(funcDecl.ReturnType, "You should change your return type to: %s", newRt).
					Push()

				// NewErrorAtNode(node.Value, "You cannot return 'none', without your value being an option type.")
				// NewErrorAtNode(funcDecl.ReturnType, "You should change your return type to: %s", newRt)
			}

		default:
			if funcDecl.ReturnType.Type != v.TypeName() {
				NewErrorAtNode(node, "[VisitReturnStatement-Literal]: Type '%s' does not match defined type of '%s'", v.TypeName(), funcDecl.ReturnType.Type)
			}

		}

	case *ast.CallExpression:
		if t := Inference.InferExpressionType(v); t != nil {
			if funcDecl.ReturnType.Type != t.TypeName() {
				NewErrorAtNode(node, "[VisitReturnStatement-CallExpression]: Type '%s' does not match defined type of '%s'", t.TypeName(), funcDecl.ReturnType.Type)
			}
		} else {
			NewErrorAtNode(node, "[VisitReturnStatement-CallExpression]: Failed to infer type of '%s'", v.GetToken())
		}

	default:
		NewErrorAtNode(node, "[VisitReturnStatement-default]: Failed to resolve type of '%s'", node.GetToken())
	}
}
func (self *TypeCheckerInstance) typeCheckIndexAccessExpression(node *ast.IndexAccessExpression) {
	// log.Debugf("TypeCheckingVisitor.VisitIndexAccessExpression: %#v", node)

	self.Scope.LinkNodeScope(node)

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

func (self *TypeCheckerInstance) typeCheckFieldAccessExpression(node *ast.FieldAccessExpression) {

	instanceType, instanceNode := TypeChecker.FindDeclaration(node)
	if instanceType == nil {
		NewErrorAtNode(node, "[VisitFieldAccessExpression]: Failed to find associated variable instance for '%s'", node.StructInstance.GetToken().String())
	}
	if instanceNode == nil {
		// NewErrorAtNode(node, "[VisitFieldAccessExpression]: Field '%s' does not exist on type '%s'", node.FieldName, instanceType.TypeName())
		NewDiagnosticAtNode(node, diagnostics.ObjectFieldNotDefined, node.FieldName, instanceType.TypeName())
	}

}
