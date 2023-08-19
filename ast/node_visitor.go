package ast

// func (self *Program) Accept(visitor NodeVisitor) any {
// 	values := make([]any, 0)
// 	for _, statement := range self.Statements {
// 		if v := statement.Accept(visitor); v != nil {
// 			values = append(values, v)
// 		}
// 	}
// 	return values
// }
// func (self *Block) Accept(visitor NodeVisitor) any {
// 	values := make([]any, 0)
// 	for _, statement := range self.Statements {
// 		if v := statement.Accept(visitor); v != nil {
// 			values = append(values, v)
// 		}
// 	}
// 	return values
// }
// func (self *Identifier) Accept(visitor NodeVisitor) any {
// 	return nil
// }
// func (self *TypedIdentifier) Accept(visitor NodeVisitor) any {
// 	if self.Identifier != nil {
// 		return self.Identifier.Accept(visitor)
// 	}
// 	return nil
// }
// func (self *ObjectDeclaration) Accept(visitor NodeVisitor) any {
// 	if self.Fields != nil {
// 		for _, field := range self.Fields {
// 			field.Accept(visitor)
// 		}
// 	}
// 	if self.Methods != nil {
// 		for _, method := range self.Methods {
// 			method.Accept(visitor)
// 		}
// 	}
// 	return nil
// }
// func (self *FunctionDeclaration) Accept(visitor NodeVisitor) any {
// 	if self.Args != nil {
// 		for _, arg := range self.Args {
// 			arg.Accept(visitor)
// 		}
// 	}
// 	if self.Receiver != nil {
// 		self.Receiver.Accept(visitor)
// 	}
// 	if self.Body != nil {
// 		self.Body.Accept(visitor)
// 	}
// 	return nil
// }
// func (self *RangeExpression) Accept(visitor NodeVisitor) any {
// 	if self.Left != nil {
// 		self.Left.Accept(visitor)
// 	}
// 	if self.Right != nil {
// 		self.Right.Accept(visitor)
// 	}
// 	return nil
// }
// func (self *AssignmentExpression) Accept(visitor NodeVisitor) any {
// 	if self.Left != nil {
// 		self.Left.Accept(visitor)
// 	}
// 	if self.Value != nil {
// 		self.Value.Accept(visitor)
// 	}
// 	return nil
// }
// func (self *BinaryExpression) Accept(visitor NodeVisitor) any {
// 	if self.Left != nil {
// 		self.Left.Accept(visitor)
// 	}
// 	if self.Right != nil {
// 		self.Right.Accept(visitor)
// 	}
// 	return nil
// }
// func (self *PostfixExpression) Accept(visitor NodeVisitor) any {
// 	if self.Left != nil {
// 		self.Left.Accept(visitor)
// 	}
// 	return nil
// }
// func (self *UnaryExpression) Accept(visitor NodeVisitor) any {
// 	if self.Expr != nil {
// 		self.Expr.Accept(visitor)
// 	}
// 	return nil
// }
// func (self *FieldAccessExpression) Accept(visitor NodeVisitor) any {
// 	if self.StructInstance != nil {
// 		self.StructInstance.Accept(visitor)
// 	}
// 	return nil
// }
// func (self *IndexAccessExpression) Accept(visitor NodeVisitor) any {
// 	if self.Instance != nil {
// 		self.Instance.Accept(visitor)
// 	}
// 	if self.StartIndex != nil {
// 		self.StartIndex.Accept(visitor)
// 	}
// 	if self.EndIndex != nil {
// 		self.EndIndex.Accept(visitor)
// 	}
// 	return nil
// }
// func (self *CallExpression) Accept(visitor NodeVisitor) any {
// 	if self.Receiver != nil {
// 		self.Receiver.Accept(visitor)
// 	}
// 	if self.Args != nil {
// 		for _, arg := range self.Args {
// 			arg.Accept(visitor)
// 		}
// 	}
// 	return nil
// }
// func (self *HttpRouteDeclaration) Accept(visitor NodeVisitor) any {
// 	if self.Path != nil {
// 		self.Path.Accept(visitor)
// 	}
// 	if self.Body != nil {
// 		self.Body.Accept(visitor)
// 	}
// 	if self.Injections != nil {
// 		for _, injection := range self.Injections {
// 			if injection != nil {
// 				injection.Accept(visitor)
// 			}
// 		}
// 	}
// 	return nil
// }
// func (self *HttpServerConfig) Accept(visitor NodeVisitor) any {
// 	if self.Port != nil {
// 		self.Port.Accept(visitor)
// 	}
// 	if self.FormMaxMemory != nil {
// 		self.FormMaxMemory.Accept(visitor)
// 	}
// 	return nil
// }
// func (self *HttpResponseData) Accept(visitor NodeVisitor) any {
// 	if self.ResponseCode != nil {
// 		self.ResponseCode.Accept(visitor)
// 	}
// 	if self.Data != nil {
// 		self.Data.Accept(visitor)
// 	}
// 	return nil
// }
// func (self *HttpRouteBodyInjection) Accept(visitor NodeVisitor) any {
// 	if self.Var != nil {
// 		self.Var.Accept(visitor)
// 	}
// 	return nil
// }
// func (self *Literal) Accept(visitor NodeVisitor) any {
//
// 	return nil
// }
// func (self *ArrayInstantiation) Accept(visitor NodeVisitor) any {
// 	if self.Type != nil {
// 		self.Type.Accept(visitor)
// 	}
// 	if self.Values != nil {
// 		for _, value := range self.Values {
// 			value.Accept(visitor)
// 		}
// 	}
// 	return nil
// }
// func (self *ObjectInstantiation) Accept(visitor NodeVisitor) any {
// 	if self.Fields != nil {
// 		for _, field := range self.Fields {
// 			if field != nil {
// 				field.Accept(visitor)
// 			}
// 		}
// 	}
// 	return nil
// }
// func (self *DictionaryInstantiation) Accept(visitor NodeVisitor) any {
// 	if self.Fields != nil {
// 		for _, field := range self.Fields {
// 			if field != nil {
// 				field.Accept(visitor)
// 			}
// 		}
// 	}
// 	return nil
// }
// func (self *IfStatement) Accept(visitor NodeVisitor) any {
// 	if self.Condition != nil {
// 		self.Condition.Accept(visitor)
// 	}
// 	if self.Body != nil {
// 		self.Body.Accept(visitor)
// 	}
// 	if self.Else != nil {
// 		self.Else.Accept(visitor)
// 	}
// 	return nil
// }
// func (self *LoopStatement) Accept(visitor NodeVisitor) any {
// 	if self.Range != nil {
// 		self.Range.Accept(visitor)
// 	}
// 	if self.Body != nil {
// 		self.Range.Accept(visitor)
// 	}
// 	if self.Step != nil {
// 		self.Range.Accept(visitor)
// 	}
// 	if self.As != nil {
// 		self.Range.Accept(visitor)
// 	}
// 	return nil
// }
// func (self *AssignmentStatement) Accept(visitor NodeVisitor) any {
// 	if self.TypedIdentifier != nil {
// 		self.TypedIdentifier.Accept(visitor)
// 	}
// 	if self.Value != nil {
// 		self.Value.Accept(visitor)
// 	}
// 	return nil
// }
// func (self *VarReference) Accept(visitor NodeVisitor) any {
//
// 	return nil
// }
// func (self *ReturnStatement) Accept(visitor NodeVisitor) any {
// 	if self.Value != nil {
// 		self.Value.Accept(visitor)
// 	}
// 	return nil
// }
// func (self *BreakStatement) Accept(visitor NodeVisitor) any {
//
// 	return nil
// }
// func (self *DeleteStatement) Accept(visitor NodeVisitor) any {
// 	if self.What != nil {
// 		self.What.Accept(visitor)
// 	}
// 	return nil
// }

//
//
//
//
//
//
//
//

func (self *Program) Accept(visitor NodeVisitor) {
	visitor.VisitProgram(self)

	for _, statement := range self.Statements {
		visitor.Visit(statement)
	}
}
func (self *Block) Accept(visitor NodeVisitor) {
	visitor.VisitBlock(self)
}
func (self *Identifier) Accept(visitor NodeVisitor) {
	visitor.VisitIdentifier(self)
}
func (self *TypedIdentifier) Accept(visitor NodeVisitor) {
	visitor.VisitTypedIdentifier(self)
	// if self.Identifier != nil {
	// 	self.Identifier.Accept(visitor)
	// }
}

func (self *TypeReference) Accept(visitor NodeVisitor) {
	visitor.VisitTypeReference(self)
}

func (self *ObjectDeclaration) Accept(visitor NodeVisitor) {
	visitor.VisitObjectDeclaration(self)
	if self.Fields != nil {
		for _, field := range self.Fields {
			visitor.VisitTypedIdentifier(field)
		}
	}
	if self.Methods != nil {
		for _, method := range self.Methods {
			visitor.VisitFunctionDeclaration(method)
		}
	}
}
func (self *FunctionDeclaration) Accept(visitor NodeVisitor) {
	// visitor.Visit(self)
	visitor.VisitFunctionDeclaration(self)
	// if self.Args != nil {
	// 	for _, arg := range self.Args {
	// 		visitor.VisitTypedIdentifier(arg)
	// 	}
	// }
	// if self.Receiver != nil {
	// 	visitor.VisitTypedIdentifier(self.Receiver)
	// }
	// if self.Body != nil {
	// 	self.Body.Accept(visitor)
	// }
	// if self.ReturnType != nil {
	// 	visitor.VisitIdentifier(self.ReturnType)
	// }
}
func (self *RangeExpression) Accept(visitor NodeVisitor) {
	// visitor.Visit(self)
	visitor.VisitRangeExpression(self)

	if self.Left != nil {
		visitor.Visit(self.Left)
	}
	if self.Right != nil {
		visitor.Visit(self.Right)
	}
}
func (self *AssignmentExpression) Accept(visitor NodeVisitor) {
	// visitor.Visit(self)
	visitor.VisitAssignmentExpression(self)
	if self.Left != nil {
		visitor.Visit(self.Left)
	}
	if self.Value != nil {
		visitor.Visit(self.Value)
	}
}
func (self *BinaryExpression) Accept(visitor NodeVisitor) {
	// visitor.Visit(self)
	visitor.VisitBinaryExpression(self)
	if self.Left != nil {
		visitor.Visit(self.Left)
	}
	if self.Right != nil {
		visitor.Visit(self.Right)
	}
}
func (self *PostfixExpression) Accept(visitor NodeVisitor) {
	// visitor.Visit(self)
	visitor.VisitPostfixExpression(self)
	if self.Left != nil {
		visitor.Visit(self.Left)
	}
}
func (self *UnaryExpression) Accept(visitor NodeVisitor) {
	// visitor.Visit(self)
	visitor.VisitUnaryExpression(self)
	if self.Expr != nil {
		visitor.Visit(self.Expr)
	}
}
func (self *FieldAccessExpression) Accept(visitor NodeVisitor) {
	// visitor.Visit(self)
	visitor.VisitFieldAccessExpression(self)
	if self.StructInstance != nil {
		visitor.Visit(self.StructInstance)
	}
}
func (self *IndexAccessExpression) Accept(visitor NodeVisitor) {
	// visitor.Visit(self)
	visitor.VisitArrayAccessExpression(self)
	if self.Instance != nil {
		visitor.Visit(self.Instance)
	}
	if self.StartIndex != nil {
		visitor.Visit(self.StartIndex)
	}
	if self.EndIndex != nil {
		visitor.Visit(self.EndIndex)
	}
}
func (self *CallExpression) Accept(visitor NodeVisitor) {
	visitor.VisitCallExpression(self)

	if self.Receiver != nil {
		self.Receiver.Accept(visitor)
	}
	if self.Args != nil {
		for _, arg := range self.Args {
			arg.Accept(visitor)
		}
	}
}
func (self *HttpRouteDeclaration) Accept(visitor NodeVisitor) {
	// visitor.Visit(self)
	visitor.VisitHttpRouteDeclaration(self)
	// if self.Path != nil {
	// 	visitor.Visit(self.Path)
	// }
	// if self.Body != nil {
	// 	self.Body.Accept(visitor)
	// }
	// if self.Injections != nil {
	// 	for _, injection := range self.Injections {
	// 		if injection != nil {
	// 			visitor.VisitHttpRouteBodyInjection(injection)
	// 		}
	// 	}
	// }
}
func (self *HttpServerConfig) Accept(visitor NodeVisitor) {
	// visitor.Visit(self)
	visitor.VisitHttpServerConfig(self)

	if self.Port != nil {
		visitor.Visit(self.Port)
	}
	if self.FormMaxMemory != nil {
		visitor.Visit(self.FormMaxMemory)
	}
}
func (self *HttpResponseData) Accept(visitor NodeVisitor) {
	// visitor.Visit(self)
	visitor.VisitHttpResponseData(self)
	if self.ResponseCode != nil {
		visitor.VisitLiteral(self.ResponseCode)
	}
	if self.Data != nil {
		visitor.Visit(self.Data)
	}
}
func (self *HttpRouteBodyInjection) Accept(visitor NodeVisitor) {
	// visitor.Visit(self)
	visitor.VisitHttpRouteBodyInjection(self)
	if self.Var != nil {
		visitor.VisitTypedIdentifier(self.Var)
	}
}
func (self *Literal) Accept(visitor NodeVisitor) {
	visitor.VisitLiteral(self)
}
func (self *ArrayInstantiation) Accept(visitor NodeVisitor) {
	// visitor.Visit(self)
	visitor.VisitArrayInstantiation(self)
	if self.Type != nil {
		visitor.VisitTypedIdentifier(self.Type)
	}
	if self.Values != nil {
		for _, value := range self.Values {
			visitor.Visit(value)
		}
	}
}
func (self *ObjectInstantiation) Accept(visitor NodeVisitor) {
	// visitor.Visit(self)
	visitor.VisitObjectInstantiation(self)

	if self.Fields != nil {
		for _, field := range self.Fields {
			if field != nil {
				visitor.Visit(field)
			}
		}
	}
}
func (self *DictionaryInstantiation) Accept(visitor NodeVisitor) {
	// visitor.Visit(self)
	visitor.VisitDictionaryInstantiation(self)
	if self.Fields != nil {
		for _, field := range self.Fields {
			if field != nil {
				visitor.Visit(field)
			}
		}
	}
}
func (self *IfStatement) Accept(visitor NodeVisitor) {
	// visitor.Visit(self)
	visitor.VisitIfStatement(self)

	if self.Condition != nil {
		visitor.Visit(self.Condition)
	}
	if self.Body != nil {
		visitor.VisitBlock(self.Body)
	}
	if self.Else != nil {
		visitor.Visit(self.Else)
	}
}
func (self *LoopStatement) Accept(visitor NodeVisitor) {
	// visitor.Visit(self)
	visitor.VisitLoopStatement(self)
	if self.Range != nil {
		visitor.Visit(self.Range)
	}
	if self.Body != nil {
		visitor.VisitBlock(self.Body)
	}
	if self.Step != nil {
		visitor.Visit(self.Range)
	}
	if self.As != nil {
		visitor.VisitIdentifier(self.As)
	}
}
func (self *AssignmentStatement) Accept(visitor NodeVisitor) {
	// visitor.Visit(self)
	visitor.VisitAssignmentStatement(self)
	if self.Type != nil {
		self.Type.Accept(visitor)
	}
	if self.Value != nil {
		self.Value.Accept(visitor)
	}
}
func (self *VarReference) Accept(visitor NodeVisitor) {
	visitor.VisitVarReference(self)
}
func (self *ReturnStatement) Accept(visitor NodeVisitor) {
	// visitor.Visit(self)
	visitor.VisitReturnStatement(self)
	if self.Value != nil {
		visitor.Visit(self.Value)
	}
}
func (self *BreakStatement) Accept(visitor NodeVisitor) {
	// visitor.Visit(self)
	visitor.VisitBreakStatement(self)
}
func (self *DeleteStatement) Accept(visitor NodeVisitor) {
	// visitor.Visit(self)
	visitor.VisitDeleteStatement(self)
	if self.What != nil {
		visitor.Visit(self.What)
	}
}

type NodeVisitor interface {
	Visit(node Node)
	VisitProgram(node *Program)
	VisitBlock(node *Block)
	VisitIdentifier(node *Identifier)
	VisitTypedIdentifier(node *TypedIdentifier)
	VisitTypeReference(node *TypeReference)
	VisitObjectDeclaration(node *ObjectDeclaration)
	VisitFunctionDeclaration(node *FunctionDeclaration)
	VisitRangeExpression(node *RangeExpression)
	VisitAssignmentExpression(node *AssignmentExpression)
	VisitBinaryExpression(node *BinaryExpression)
	VisitPostfixExpression(node *PostfixExpression)
	VisitUnaryExpression(node *UnaryExpression)
	VisitFieldAccessExpression(node *FieldAccessExpression)
	VisitArrayAccessExpression(node *IndexAccessExpression)
	VisitCallExpression(node *CallExpression)
	VisitHttpRouteDeclaration(node *HttpRouteDeclaration)
	VisitHttpServerConfig(node *HttpServerConfig)
	VisitHttpResponseData(node *HttpResponseData)
	VisitHttpRouteBodyInjection(node *HttpRouteBodyInjection)
	VisitLiteral(node *Literal)
	VisitArrayInstantiation(node *ArrayInstantiation)
	VisitObjectInstantiation(node *ObjectInstantiation)
	VisitDictionaryInstantiation(node *DictionaryInstantiation)
	VisitIfStatement(node *IfStatement)
	VisitLoopStatement(node *LoopStatement)
	VisitAssignmentStatement(node *AssignmentStatement)
	VisitVarReference(node *VarReference)
	VisitReturnStatement(node *ReturnStatement)
	VisitBreakStatement(node *BreakStatement)
	VisitDeleteStatement(node *DeleteStatement)
}

type NodeVisitorAdapter struct{}

func (self *NodeVisitorAdapter) Visit(node Node) {
	switch node := node.(type) {
	case *Program:
		// self.VisitProgram(node)
		node.Accept(self)
	case *Block:
		// self.VisitBlock(node)
		node.Accept(self)
	case *Identifier:
		// self.VisitIdentifier(node)
		node.Accept(self)
	case *TypeReference:
		node.Accept(self)
	case *TypedIdentifier:
		// self.VisitTypedIdentifier(node)
		node.Accept(self)
	case *ObjectDeclaration:
		// self.VisitObjectDeclaration(node)
		node.Accept(self)
	case *FunctionDeclaration:
		// self.VisitFunctionDeclaration(node)
		node.Accept(self)
	case *RangeExpression:
		// self.VisitRangeExpression(node)
		node.Accept(self)
	case *AssignmentExpression:
		// self.VisitAssignmentExpression(node)
		node.Accept(self)
	case *BinaryExpression:
		// self.VisitBinaryExpression(node)
		node.Accept(self)
	case *PostfixExpression:
		// self.VisitPostfixExpression(node)
		node.Accept(self)
	case *UnaryExpression:
		// self.VisitUnaryExpression(node)
		node.Accept(self)
	case *FieldAccessExpression:
		// self.VisitFieldAccessExpression(node)
		node.Accept(self)
	case *IndexAccessExpression:
		// self.VisitArrayAccessExpression(node)
		node.Accept(self)
	case *CallExpression:
		// self.VisitCallExpression(node)
		node.Accept(self)
	case *HttpRouteDeclaration:
		// self.VisitHttpRouteDeclaration(node)
		node.Accept(self)
	case *HttpServerConfig:
		// self.VisitHttpServerConfig(node)
		node.Accept(self)
	case *HttpResponseData:
		// self.VisitHttpResponseData(node)
		node.Accept(self)
	case *HttpRouteBodyInjection:
		// self.VisitHttpRouteBodyInjection(node)
		node.Accept(self)
	case *Literal:
		// self.VisitLiteral(node)
		node.Accept(self)
	case *ArrayInstantiation:
		// self.VisitArrayInstantiation(node)
		node.Accept(self)
	case *ObjectInstantiation:
		// self.VisitObjectInstantiation(node)
		node.Accept(self)
	case *DictionaryInstantiation:
		// self.VisitDictionaryInstantiation(node)
		node.Accept(self)
	case *IfStatement:
		// self.VisitIfStatement(node)
		node.Accept(self)
	case *LoopStatement:
		// self.VisitLoopStatement(node)
		node.Accept(self)
	case *AssignmentStatement:
		// self.VisitAssignmentStatement(node)
		node.Accept(self)
	case *VarReference:
		// self.VisitVarReference(node)
		node.Accept(self)
	case *ReturnStatement:
		// self.VisitReturnStatement(node)
		node.Accept(self)
	case *BreakStatement:
		// self.VisitBreakStatement(node)
		node.Accept(self)
	case *DeleteStatement:
		// self.VisitDeleteStatement(node)
		node.Accept(self)
	}
}

func (self *NodeVisitorAdapter) VisitProgram(node *Program) {

}

func (self *NodeVisitorAdapter) VisitBlock(node *Block) {

}

func (self *NodeVisitorAdapter) VisitIdentifier(node *Identifier) {

}

func (self *NodeVisitorAdapter) VisitTypeReference(node *TypeReference) {

}
func (self *NodeVisitorAdapter) VisitTypedIdentifier(node *TypedIdentifier) {

}

func (self *NodeVisitorAdapter) VisitObjectDeclaration(node *ObjectDeclaration) {

}

func (self *NodeVisitorAdapter) VisitFunctionDeclaration(node *FunctionDeclaration) {

}

func (self *NodeVisitorAdapter) VisitRangeExpression(node *RangeExpression) {

}

func (self *NodeVisitorAdapter) VisitAssignmentExpression(node *AssignmentExpression) {

}

func (self *NodeVisitorAdapter) VisitBinaryExpression(node *BinaryExpression) {

}

func (self *NodeVisitorAdapter) VisitPostfixExpression(node *PostfixExpression) {

}

func (self *NodeVisitorAdapter) VisitUnaryExpression(node *UnaryExpression) {

}

func (self *NodeVisitorAdapter) VisitFieldAccessExpression(node *FieldAccessExpression) {

}

func (self *NodeVisitorAdapter) VisitArrayAccessExpression(node *IndexAccessExpression) {

}

func (self *NodeVisitorAdapter) VisitCallExpression(node *CallExpression) {

}

func (self *NodeVisitorAdapter) VisitHttpRouteDeclaration(node *HttpRouteDeclaration) {

}

func (self *NodeVisitorAdapter) VisitHttpServerConfig(node *HttpServerConfig) {

}

func (self *NodeVisitorAdapter) VisitHttpResponseData(node *HttpResponseData) {

}

func (self *NodeVisitorAdapter) VisitHttpRouteBodyInjection(node *HttpRouteBodyInjection) {

}

func (self *NodeVisitorAdapter) VisitLiteral(node *Literal) {

}

func (self *NodeVisitorAdapter) VisitArrayInstantiation(node *ArrayInstantiation) {

}

func (self *NodeVisitorAdapter) VisitObjectInstantiation(node *ObjectInstantiation) {

}

func (self *NodeVisitorAdapter) VisitDictionaryInstantiation(node *DictionaryInstantiation) {

}

func (self *NodeVisitorAdapter) VisitIfStatement(node *IfStatement) {

}

func (self *NodeVisitorAdapter) VisitLoopStatement(node *LoopStatement) {

}

func (self *NodeVisitorAdapter) VisitAssignmentStatement(node *AssignmentStatement) {

}

func (self *NodeVisitorAdapter) VisitVarReference(node *VarReference) {

}

func (self *NodeVisitorAdapter) VisitReturnStatement(node *ReturnStatement) {

}

func (self *NodeVisitorAdapter) VisitBreakStatement(node *BreakStatement) {

}

func (self *NodeVisitorAdapter) VisitDeleteStatement(node *DeleteStatement) {

}
