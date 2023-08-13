package ast

type Type interface {
	TypeName() string
}

type BasicType struct {
	Name string
}

func (self *BasicType) TypeName() string {
	return self.Name
}

var (
	IntType    = &BasicType{"int"}
	StringType = &BasicType{"string"}
	BoolType   = &BasicType{"bool"}
	FloatType  = &BasicType{"float"}
	NullType   = &BasicType{"null"}
	VoidType   = &BasicType{"void"}
)

var BasicTypes = map[string]Type{
	"int":    IntType,
	"string": StringType,
	"bool":   BoolType,
	"float":  FloatType,
	"null":   NullType,
	"void":   VoidType,
}
