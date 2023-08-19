package ast

type ArrayInstantiation struct {
	*AstNode

	Type   *TypedIdentifier
	Values []Expr
}

func (self *ArrayInstantiation) GetChildren() []Node {
	var nodes []Node
	if self.Type != nil {
		nodes = append(nodes, self.Type)
	}
	for _, value := range self.Values {
		nodes = append(nodes, value)
	}
	return nodes
}
func (self *ArrayInstantiation) IsExpression() {}
func (self *ArrayInstantiation) IsStatement()  {}

type ObjectInstantiation struct {
	*AstNode

	TypeName *Identifier
	Fields   map[string]Expr
}

func (self *ObjectInstantiation) GetChildren() []Node {
	var nodes []Node
	nodes = append(nodes, self.TypeName)
	for _, field := range self.Fields {
		nodes = append(nodes, field)
	}
	return nodes
}
func (self *ObjectInstantiation) IsExpression() {}

type DictionaryInstantiation struct {
	*AstNode
	Fields map[string]Expr
}

func (self *DictionaryInstantiation) GetChildren() []Node {
	var nodes []Node
	for _, field := range self.Fields {
		nodes = append(nodes, field)
	}
	return nodes
}
func (self *DictionaryInstantiation) IsExpression() {}
