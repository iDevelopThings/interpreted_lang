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

func (self *BasicType) TypeName() string          { return self.Name }
func (self *BasicType) GetEnvBindingName() string { return self.Name }

func (self *BasicType) IsCompatibleWith(typ Type) bool {
	if self.TypeName() == typ.TypeName() {
		return true
	}

	if t, ok := typ.(*Literal); ok {
		typ = t.GetBasicType()
	}

	if t, ok := typ.(*BasicType); ok {
		if t.TypeName() == "any" {
			return true
		}

		return true
	}

	// switch {
	//
	// case self.TypeName() == "string":
	// 	{
	// 		switch typ.TypeName() {
	// 		case "int", "float", "bool":
	// 			return true
	// 		default:
	// 			return false
	// 		}
	// 	}
	//
	// case self.TypeName() == "int":
	// 	{
	// 		switch typ.TypeName() {
	// 		case "float", "bool", "string":
	// 			return true
	// 		default:
	// 			return false
	// 		}
	// 	}
	//
	// case self.TypeName() == "float":
	// 	{
	//
	// 	}

	return false
}

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
