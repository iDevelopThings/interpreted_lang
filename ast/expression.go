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

func (self *BinaryExpression) GetChildren() []Node {
	var result []Node
	result = append(result, self.Left)
	result = append(result, self.Right)
	return result
}
func (self *BinaryExpression) IsExpression() {}
func (self *BinaryExpression) IsStatement()  {}

type UnaryExpression struct {
	*AstNode
	Op   operators.Operator `visitor:"-"`
	Left Expr
}

func (self *UnaryExpression) GetChildren() []Node {
	var result []Node
	result = append(result, self.Left)
	return result
}
func (self *UnaryExpression) IsExpression() {}
func (self *UnaryExpression) IsStatement()  {}

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
func (self *FieldAccessExpression) IsStatement()  {}

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

func (self *RangeExpression) GetChildren() []Node {
	var result []Node
	result = append(result, self.Left)
	result = append(result, self.Right)
	return result
}
func (self *RangeExpression) IsExpression() {}
