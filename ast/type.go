package ast

type Type interface {
	Node
	TypeName() string
	GetEnvBindingName() string
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

func (self *BasicType) TypeName() string          { return self.Name }
func (self *BasicType) GetEnvBindingName() string { return self.Name }

var (
	IntType = &BasicType{
		AstNode: NewAstNode(nil),
		Name:    "int",
	}
	StringType = &BasicType{
		AstNode: NewAstNode(nil),
		Name:    "string",
	}
	BoolType = &BasicType{
		AstNode: NewAstNode(nil),
		Name:    "bool",
	}
	FloatType = &BasicType{
		AstNode: NewAstNode(nil),
		Name:    "float",
	}
	NoneType = &BasicType{
		AstNode: NewAstNode(nil),
		Name:    "none",
	}
	VoidType = &BasicType{
		AstNode: NewAstNode(nil),
		Name:    "void",
	}
	AnyType = &BasicType{
		AstNode: NewAstNode(nil),
		Name:    "any",
	}
)

var BasicTypes = map[string]Type{
	"int":    IntType,
	"string": StringType,
	"bool":   BoolType,
	"float":  FloatType,
	"none":   NoneType,
	"void":   VoidType,
	"any":    AnyType,
}
