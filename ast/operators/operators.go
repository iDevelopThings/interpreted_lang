package operators

import (
	"github.com/charmbracelet/log"
)

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
	BitwiseXor         Operator = "^"
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
	BitwiseXor,
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

// Note that in Go, interfaces are satisfied implicitly.
// Thus, any type which is a int or float64 will satisfy BinaryOperatable.
type BinaryOperatable interface {
	~int | ~float64
}

// func BinaryOperation[T BinaryOperatable](op Operator, a, b T) T {
// 	switch v := a.(type) {
// 	case int:
// 		// We can safely assert b is also of type int because of type constraints.
// 		return BinaryIntOperation(op, v, b.(int)).(T)
// 	case float64:
// 		// We can safely assert b is also of type float64 because of type constraints.
// 		return BinaryFloatOperation(op, v, b.(float64)).(T)
// 	}
// 	panic("Unsupported operation: " + string(op))
// }

func BinaryIntComparisonOperation(op Operator, a, b int) bool {
	switch op {
	case EqualEqual:
		return a == b
	case NotEqual:
		return a != b
	case LessThan:
		return a < b
	case GreaterThan:
		return a > b
	case LessThanOrEqual:
		return a <= b
	case GreaterThanOrEqual:
		return a >= b
	default:
		log.Fatalf("Unsupported int comparison operation: lhs=%d, rhs=%d, op=%s", a, b, string(op))
	}
	return false
}

func BinaryIntOperation(op Operator, a, b int) int {
	switch op {
	case PlusPlus:
		return a + 1
	case MinusMinus:
		return a - 1
	case Plus, PlusEqual:
		return a + b
	case Minus, MinusEqual:
		return a - b
	case Multiply, MultiplyEqual:
		return a * b
	case Divide, DivideEqual:
		if b == 0 {
			log.Fatalf("Division by zero: lhs=%d, rhs=%d", a, b)
		}
		return a / b
	default:
		log.Fatalf("Unsupported int operation: lhs=%d, rhs=%d, op=%s", a, b, string(op))
	}
	return 0
}

func BinaryFloatComparisonOperation(op Operator, a, b float64) bool {
	switch op {
	case EqualEqual:
		return a == b
	case NotEqual:
		return a != b
	case LessThan:
		return a < b
	case GreaterThan:
		return a > b
	case LessThanOrEqual:
		return a <= b
	case GreaterThanOrEqual:
		return a >= b
	default:
		log.Fatalf("Unsupported int comparison operation: lhs=%d, rhs=%d, op=%s", a, b, string(op))
	}
	return false
}

func BinaryFloatOperation(op Operator, a, b float64) float64 {
	switch op {
	case PlusPlus:
		return a + 1.0
	case MinusMinus:
		return a - 1.0
	case Plus, PlusEqual:
		return a + b
	case Minus, MinusEqual:
		return a - b
	case Multiply, MultiplyEqual:
		return a * b
	case Divide, DivideEqual:
		if b == 0.0 {
			log.Fatalf("Division by zero: lhs=%f, rhs=%f", a, b)
		}
		return a / b
	default:
		log.Fatalf("Unsupported float operation: lhs=%f, rhs=%f, op=%s", a, b, string(op))
		return 0.0
	}
}
