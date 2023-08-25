package operators

type Operator string

const (
	Plus               Operator = "+"
	PlusPlus           Operator = "++"
	PlusEqual          Operator = "+="
	Minus              Operator = "-"
	MinusMinus         Operator = "--"
	MinusEqual         Operator = "-="
	Multiply           Operator = "*"
	MultiplyEqual      Operator = "*="
	Divide             Operator = "/"
	DivideEqual        Operator = "/="
	Modulo             Operator = "%"
	Power              Operator = "^"
	And                Operator = "&&"
	Or                 Operator = "||"
	Equal              Operator = "="
	EqualEqual         Operator = "=="
	NotEqual           Operator = "!="
	LessThan           Operator = "<"
	GreaterThan        Operator = ">"
	LessThanOrEqual    Operator = "<="
	GreaterThanOrEqual Operator = ">="
	BitwiseAnd         Operator = "&"
	BitwiseOr          Operator = "|"
	BitwiseNot         Operator = "~"
	BitwiseLeftShift   Operator = "<<"
	BitwiseRightShift  Operator = ">>"
)

var AllOperators = []Operator{
	Plus,
	PlusPlus,
	PlusEqual,
	Minus,
	MinusMinus,
	MinusEqual,
	Multiply,
	MultiplyEqual,
	Divide,
	DivideEqual,
	Modulo,
	Power,
	And,
	Or,
	Equal,
	EqualEqual,
	NotEqual,
	LessThan,
	GreaterThan,
	LessThanOrEqual,
	GreaterThanOrEqual,
	BitwiseAnd,
	BitwiseOr,
	BitwiseNot,
	BitwiseLeftShift,
	BitwiseRightShift,
}

var OperatorMap = map[string]Operator{}

func ToOperator(value string) Operator {
	return OperatorMap[value]
}

func init() {
	for _, op := range AllOperators {
		OperatorMap[string(op)] = op
	}
}
