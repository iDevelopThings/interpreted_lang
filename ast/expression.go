package ast

import (
	"arc/ast/operators"
)

type BinaryExpressionKind string

const (
	BinaryExpressionKindUnknown    BinaryExpressionKind = "Unknown"
	BinaryExpressionKindRegular    BinaryExpressionKind = "Regular"
	BinaryExpressionKindAssignment BinaryExpressionKind = "Assignment"
	BinaryExpressionKindComparison BinaryExpressionKind = "Comparison"
)

type BinaryExpression struct {
	*AstNode
	Kind  BinaryExpressionKind
	Left  Expr
	Op    operators.Operator `visitor:"-"`
	Right Expr
}

func (self *BinaryExpression) IsExpression() {}
func (self *BinaryExpression) IsStatement()  {}

type UnaryExpression struct {
	*AstNode
	Op   operators.Operator `visitor:"-"`
	Left Expr
}

func (self *UnaryExpression) IsExpression() {}
func (self *UnaryExpression) IsStatement()  {}

type FieldAccessExpression struct {
	*AstNode
	StructInstance Expr
	FieldName      string
	StaticAccess   bool
}

func (self *FieldAccessExpression) IsExpression() {}
func (self *FieldAccessExpression) IsStatement()  {}

type IndexAccessExpression struct {
	*AstNode
	Instance Expr

	StartIndex Expr
	EndIndex   Expr

	IsSlice bool
}

func (self *IndexAccessExpression) IsExpression() {}

type ExpressionList struct {
	*AstNode
	Entries []Expr
}

func (self *ExpressionList) IsExpression() {}

type CallExpression struct {
	*AstNode
	Function *Identifier
	Receiver Expr

	ArgumentList *ExpressionList

	IsStaticAccess bool
}

func (self *CallExpression) IsStatement()  {}
func (self *CallExpression) IsExpression() {}

// Is not an `||` expression
// this is `someFunction() or { ... }`
type OrExpression struct {
	*AstNode
	Left  Expr
	Right Expr
}

func (self *OrExpression) IsStatement()  {}
func (self *OrExpression) IsExpression() {}

type RangeExpression struct {
	*AstNode
	Left  Expr
	Right Expr
}

func (self *RangeExpression) IsExpression() {}
