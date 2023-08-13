package ast

type ImportStatement struct {
	*AstNode
	Path *Literal
}

func (self *ImportStatement) IsStatement()         {}
func (self *ImportStatement) IsTopLevelStatement() {}

type IfStatement struct {
	*AstNode
	Condition Expr
	Body      *Block
	Else      Statement // Either a Block or another IfStatement
}

func (self *IfStatement) IsStatement() {}

type LoopStatement struct {
	*AstNode
	Range Expr
	Body  *Block
	Step  Expr
	As    *Identifier
}

func (self *LoopStatement) IsStatement() {}

type AssignmentStatement struct {
	*AstNode
	*TypedIdentifier
	Value Expr
}

func (self *AssignmentStatement) IsStatement()         {}
func (self *AssignmentStatement) IsTopLevelStatement() {}

type VarReference struct {
	*AstNode
	Name string
}

func (self *VarReference) IsExpression() {}

type ReturnStatement struct {
	*AstNode
	Value Expr
}

func (self *ReturnStatement) IsStatement() {}

type BreakStatement struct {
	*AstNode
}

func (self *BreakStatement) IsStatement() {}

type DeleteStatement struct {
	*AstNode
	What Expr
}

func (self *DeleteStatement) IsStatement() {}
