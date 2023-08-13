package ast

import (
	"interpreted_lang/ast/operators"
)

type RangeExpression struct {
	*AstNode
	Left  Expr
	Right Expr
}

func (self *RangeExpression) IsExpression() {}

type AssignmentExpression struct {
	*AstNode
	Left  Expr
	Op    operators.Operator
	Value Expr
}

func (self *AssignmentExpression) IsExpression() {}
func (self *AssignmentExpression) IsStatement()  {}

type BinaryExpressionKind string

const (
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

func (self *BinaryExpression) IsExpression() {}
func (self *BinaryExpression) IsStatement()  {}

type PostfixExpression struct {
	*AstNode
	Left Expr
	Op   operators.Operator
}

func (self *PostfixExpression) IsExpression() {}
func (self *PostfixExpression) IsStatement()  {}

type UnaryExpression struct {
	*AstNode
	Op   operators.Operator
	Expr Expr
}

func (self *UnaryExpression) IsExpression() {}

type FieldAccessExpression struct {
	*AstNode
	StructInstance Expr
	FieldName      string
}

func (self *FieldAccessExpression) IsExpression() {}

type ArrayAccessExpression struct {
	*AstNode
	Instance Expr

	StartIndex Expr
	EndIndex   Expr

	IsSlice bool
}

func (self *ArrayAccessExpression) IsExpression() {}

type CallExpression struct {
	*AstNode
	FunctionName string
	Receiver     Expr
	Args         []Expr
}

func (self *CallExpression) IsStatement()  {}
func (self *CallExpression) IsExpression() {}
