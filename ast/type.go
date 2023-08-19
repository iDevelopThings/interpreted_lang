package ast

type Type interface {
	Node
	TypeName() string
}

type BasicType struct {
	*AstNode
	Name string
}

func (self *BasicType) GetChildren() []Node {
	return []Node{}
}

func (self *BasicType) Accept(visitor NodeVisitor) {
	visitor.Visit(self)
}

func (self *BasicType) TypeName() string {
	return self.Name
}

var (
	IntType = &BasicType{
		AstNode: &AstNode{NodeId: GetUniqueNodeId()},
		Name:    "int",
	}
	StringType = &BasicType{
		AstNode: &AstNode{NodeId: GetUniqueNodeId()},
		Name:    "string",
	}
	BoolType = &BasicType{
		AstNode: &AstNode{NodeId: GetUniqueNodeId()},
		Name:    "bool",
	}
	FloatType = &BasicType{
		AstNode: &AstNode{NodeId: GetUniqueNodeId()},
		Name:    "float",
	}
	NullType = &BasicType{
		AstNode: &AstNode{NodeId: GetUniqueNodeId()},
		Name:    "null",
	}
	VoidType = &BasicType{
		AstNode: &AstNode{NodeId: GetUniqueNodeId()},
		Name:    "void",
	}
	AnyType = &BasicType{
		AstNode: &AstNode{NodeId: GetUniqueNodeId()},
		Name:    "any",
	}
)

var BasicTypes = map[string]Type{
	"int":    IntType,
	"string": StringType,
	"bool":   BoolType,
	"float":  FloatType,
	"null":   NullType,
	"void":   VoidType,
	"any":    AnyType,
}
