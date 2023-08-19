package ast

type ImportStatement struct {
	*AstNode
	Path *Literal
}

func (self *ImportStatement) Accept(visitor NodeVisitor) {
	visitor.Visit(self)
}

func (self *ImportStatement) GetChildren() []Node {
	return []Node{self.Path}
}
func (self *ImportStatement) IsStatement()         {}
func (self *ImportStatement) IsTopLevelStatement() {}

type IfStatement struct {
	*AstNode
	Condition Expr
	Body      *Block
	Else      Statement // Either a Block or another IfStatement
}

func (self *IfStatement) GetChildren() []Node {
	return []Node{self.Condition, self.Body, self.Else}
}
func (self *IfStatement) IsStatement() {}

type LoopStatement struct {
	*AstNode
	Range Expr
	Body  *Block
	Step  Expr
	As    *Identifier
}

func (self *LoopStatement) GetChildren() []Node {
	return []Node{self.Range, self.Body, self.Step, self.As}
}
func (self *LoopStatement) IsStatement() {}

type AssignmentStatement struct {
	*AstNode
	// Type  *TypedIdentifier
	Name  *Identifier
	Type  *TypeReference
	Value Expr
}

func (self *AssignmentStatement) GetChildren() []Node {
	return []Node{self.Type, self.Value}
}
func (self *AssignmentStatement) IsStatement()         {}
func (self *AssignmentStatement) IsTopLevelStatement() {}

type VarReference struct {
	*AstNode
	Name string
}

func (self *VarReference) GetChildren() []Node {
	return []Node{}
}
func (self *VarReference) IsExpression() {}

type ReturnStatement struct {
	*AstNode
	Value Expr
}

func (self *ReturnStatement) GetChildren() []Node {
	return []Node{self.Value}
}
func (self *ReturnStatement) IsStatement() {}

type BreakStatement struct {
	*AstNode
}

func (self *BreakStatement) GetChildren() []Node {
	return []Node{}
}
func (self *BreakStatement) IsStatement() {}

type DeleteStatement struct {
	*AstNode
	What Expr
}

func (self *DeleteStatement) GetChildren() []Node {
	return []Node{self.What}
}
func (self *DeleteStatement) IsStatement() {}
