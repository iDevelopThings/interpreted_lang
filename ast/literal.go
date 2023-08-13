package ast

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/goccy/go-json"
)

type LiteralKind string

const (
	LiteralKindUnknown LiteralKind = "unknown"
	LiteralKindString  LiteralKind = "string"
	LiteralKindInteger LiteralKind = "int"
	LiteralKindFloat   LiteralKind = "float"
	LiteralKindBoolean LiteralKind = "bool"
	LiteralKindNull    LiteralKind = "null"
)

var AllLiteralKinds = []LiteralKind{
	LiteralKindUnknown,
	LiteralKindString,
	LiteralKindInteger,
	LiteralKindFloat,
	LiteralKindBoolean,
	LiteralKindNull,
}

type Literal struct {
	*AstNode
	Kind  LiteralKind
	Value any
}

func (self *Literal) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.Value)
}

func (self *Literal) IsExpression() {}

func NewLiteral(ctx antlr.ParserRuleContext, value any) *Literal {
	lit := &Literal{
		AstNode: NewAstNode(ctx),
	}

	switch value.(type) {
	case string:
		lit.Kind = LiteralKindString
		lit.Value = value.(string)
	case int:
		lit.Kind = LiteralKindInteger
		lit.Value = value.(int)
	case float64:
		lit.Kind = LiteralKindFloat
		lit.Value = value.(float64)
	case bool:
		lit.Kind = LiteralKindBoolean
		lit.Value = value.(bool)
	case nil:
		lit.Kind = LiteralKindNull
		lit.Value = nil
	default:
		panic("Unknown literal type: " + fmt.Sprintf("%T", value))
	}

	return lit
}

func (self *Literal) Unescape() error {
	if self.Kind != LiteralKindString {
		return nil
	}

	value := self.Value.(string)
	var result strings.Builder
	i := 0
	for i < len(value) {
		if value[i] != '\\' {
			result.WriteByte(value[i])
			i++
		} else {
			if i+1 >= len(value) {
				return fmt.Errorf("invalid escape sequence at end of string")
			}
			switch value[i+1] {
			case 'b':
				result.WriteByte('\b')
			case 'f':
				result.WriteByte('\f')
			case 'n':
				result.WriteByte('\n')
			case 'r':
				result.WriteByte('\r')
			case 't':
				result.WriteByte('\t')
			case '\\':
				result.WriteByte('\\')
			case '"':
				result.WriteByte('"')
			case '\'':
				result.WriteByte('\'')
			case 'x':
				if i+3 >= len(value) {
					return fmt.Errorf("invalid \\x escape sequence")
				}
				v, err := strconv.ParseUint(value[i+2:i+4], 16, 8)
				if err != nil {
					return fmt.Errorf("invalid \\x escape sequence")
				}
				result.WriteByte(byte(v))
				i += 3
			case 'u':
				if i+5 >= len(value) {
					return fmt.Errorf("invalid \\u escape sequence")
				}
				v, err := strconv.ParseUint(value[i+2:i+6], 16, 16)
				if err != nil {
					return fmt.Errorf("invalid \\u escape sequence")
				}
				result.WriteRune(rune(v))
				i += 5
			default:
				return fmt.Errorf("invalid escape sequence \\%c", value[i+1])
			}
			i += 2
		}
	}

	self.Value = result.String()

	return nil
}

func (self *Literal) IsGreaterThan(other *Literal) bool {
	if self.Kind != other.Kind {
		return false
	}

	switch self.Kind {
	case LiteralKindInteger:
		return self.Value.(int) > other.Value.(int)
	case LiteralKindFloat:
		return self.Value.(float64) > other.Value.(float64)
	case LiteralKindString:
		return self.Value.(string) > other.Value.(string)
	default:
		panic("Cannot compare literal")
	}
}
func (self *Literal) IsLessThan(other *Literal) bool {
	if self.Kind != other.Kind {
		return false
	}

	switch self.Kind {
	case LiteralKindInteger:
		return self.Value.(int) < other.Value.(int)
	case LiteralKindFloat:
		return self.Value.(float64) < other.Value.(float64)
	case LiteralKindString:
		return self.Value.(string) < other.Value.(string)
	default:
		panic("Cannot compare literal")
	}
}
func (self *Literal) IsGreaterThanOrEqual(other *Literal) bool {
	if self.Kind != other.Kind {
		return false
	}

	switch self.Kind {
	case LiteralKindInteger:
		return self.Value.(int) >= other.Value.(int)
	case LiteralKindFloat:
		return self.Value.(float64) >= other.Value.(float64)
	case LiteralKindString:
		return self.Value.(string) >= other.Value.(string)
	default:
		panic("Cannot compare literal")
	}
}
func (self *Literal) IsLessThanOrEqual(other *Literal) bool {
	if self.Kind != other.Kind {
		return false
	}

	switch self.Kind {
	case LiteralKindInteger:
		return self.Value.(int) <= other.Value.(int)
	case LiteralKindFloat:
		return self.Value.(float64) <= other.Value.(float64)
	case LiteralKindString:
		return self.Value.(string) <= other.Value.(string)
	default:
		panic("Cannot compare literal")
	}
}
func (self *Literal) IsEqual(other *Literal) bool {
	if self.Kind != other.Kind {
		return false
	}

	switch self.Kind {
	case LiteralKindInteger:
		return self.Value.(int) == other.Value.(int)
	case LiteralKindFloat:
		return self.Value.(float64) == other.Value.(float64)
	case LiteralKindString:
		return self.Value.(string) == other.Value.(string)
	case LiteralKindBoolean:
		return self.Value.(bool) == other.Value.(bool)
	case LiteralKindNull:
		return true
	default:
		panic("Cannot compare literal")
	}
}
func (self *Literal) IsNotEqual(other *Literal) bool {
	if self.Kind != other.Kind {
		return true
	}

	switch self.Kind {
	case LiteralKindInteger:
		return self.Value.(int) != other.Value.(int)
	case LiteralKindFloat:
		return self.Value.(float64) != other.Value.(float64)
	case LiteralKindString:
		return self.Value.(string) != other.Value.(string)
	case LiteralKindBoolean:
		return self.Value.(bool) != other.Value.(bool)
	case LiteralKindNull:
		return false
	default:
		panic("Cannot compare literal")
	}
}

// Allow an int/float or literal to be added to a literal
func (self *Literal) Add(other any) {
	switch other := other.(type) {
	case *Literal:
		if self.Kind != other.Kind {
			panic("Cannot add literals of different types")
		}

		switch self.Kind {
		case LiteralKindInteger:
			self.Value = self.Value.(int) + other.Value.(int)
		case LiteralKindFloat:
			self.Value = self.Value.(float64) + other.Value.(float64)
		case LiteralKindString:
			self.Value = self.Value.(string) + other.Value.(string)
		default:
			panic("Cannot add literals of this type")
		}
	case int:
		if self.Kind != LiteralKindInteger {
			panic("Cannot add int to non-integer literal")
		}

		self.Value = self.Value.(int) + other
	case float64:
		if self.Kind != LiteralKindFloat {
			panic("Cannot add float to non-float literal")
		}

		self.Value = self.Value.(float64) + other
	default:
		panic("Cannot add non-literal to literal")
	}

}
