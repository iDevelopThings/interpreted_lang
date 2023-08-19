package ast

import (
	"arc/ast/operators"
)

type RangeExpression struct {
	*AstNode
	Left  Expr
	Right Expr
}

func (self *RangeExpression) GetChildren() []Node {
	var result []Node
	result = append(result, self.Left)
	result = append(result, self.Right)
	return result
}
func (self *RangeExpression) IsExpression() {}

type AssignmentExpression struct {
	*AstNode
	Left  Expr
	Op    operators.Operator
	Value Expr
}

func (self *AssignmentExpression) GetChildren() []Node {
	var result []Node
	result = append(result, self.Left)
	result = append(result, self.Value)
	return result
}
func (self *AssignmentExpression) IsExpression() {}
func (self *AssignmentExpression) IsStatement()  {}

type BinaryExpressionKind string

const (
	BinaryExpressionKindUnknown        BinaryExpressionKind = "Unknown"
	BinaryExpressionKindAssignment     BinaryExpressionKind = "Assignment"
	BinaryExpressionKindMultiplicative BinaryExpressionKind = "Multiplicative"
	BinaryExpressionKindAdditive       BinaryExpressionKind = "Additive"
	BinaryExpressionKindEquality       BinaryExpressionKind = "Equality"
	BinaryExpressionKindRelational     BinaryExpressionKind = "Relational"
	BinaryExpressionKindShift          BinaryExpressionKind = "Shift"
	BinaryExpressionKindLogicalAnd     BinaryExpressionKind = "LogicalAnd"
	BinaryExpressionKindLogicalOr      BinaryExpressionKind = "LogicalOr"
	BinaryExpressionKindPower          BinaryExpressionKind = "Power"
)

type BinaryExpression struct {
	*AstNode
	Kind  BinaryExpressionKind
	Left  Expr
	Op    operators.Operator
	Right Expr
}

func (self *BinaryExpression) GetChildren() []Node {
	var result []Node
	result = append(result, self.Left)
	result = append(result, self.Right)
	return result
}
func (self *BinaryExpression) IsExpression() {}
func (self *BinaryExpression) IsStatement()  {}

type PostfixExpression struct {
	*AstNode
	Left Expr
	Op   operators.Operator
}

func (self *PostfixExpression) GetChildren() []Node {
	var result []Node
	result = append(result, self.Left)
	return result
}
func (self *PostfixExpression) IsExpression() {}
func (self *PostfixExpression) IsStatement()  {}

type UnaryExpression struct {
	*AstNode
	Op   operators.Operator
	Expr Expr
}

func (self *UnaryExpression) GetChildren() []Node {
	var result []Node
	result = append(result, self.Expr)
	return result
}
func (self *UnaryExpression) IsExpression() {}

type FieldAccessExpression struct {
	*AstNode
	StructInstance Expr
	FieldName      string
	StaticAccess   bool
}

func (self *FieldAccessExpression) GetChildren() []Node {
	var result []Node
	result = append(result, self.StructInstance)
	return result
}
func (self *FieldAccessExpression) IsExpression() {}

type IndexAccessExpression struct {
	*AstNode
	Instance Expr

	StartIndex Expr
	EndIndex   Expr

	IsSlice bool
}

func (self *IndexAccessExpression) GetChildren() []Node {
	var result []Node
	result = append(result, self.Instance)
	result = append(result, self.StartIndex)
	if self.IsSlice {
		result = append(result, self.EndIndex)
	}
	return result
}
func (self *IndexAccessExpression) IsExpression() {}

type CallExpression struct {
	*AstNode
	Function *Identifier
	Receiver Expr

	Args           []Expr
	IsStaticAccess bool
}

func (self *CallExpression) GetChildren() []Node {
	var result []Node
	if self.Receiver != nil {
		result = append(result, self.Receiver)
	}
	for _, arg := range self.Args {
		result = append(result, arg)
	}
	return result
}
func (self *CallExpression) IsStatement()  {}
func (self *CallExpression) IsExpression() {}
