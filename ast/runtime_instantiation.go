package ast

type ArrayInstantiation struct {
	*AstNode

	Type   *TypedIdentifier
	Values []Expr
}

func (self *ArrayInstantiation) IsExpression() {}
func (self *ArrayInstantiation) IsStatement()  {}

type ObjectInstantiation struct {
	*AstNode

	TypeName *Identifier
	Fields   map[string]Expr
}

func (self *ObjectInstantiation) IsExpression() {}

type DictionaryInstantiation struct {
	*AstNode
	Fields map[string]Expr
}

func (self *DictionaryInstantiation) IsExpression() {}
