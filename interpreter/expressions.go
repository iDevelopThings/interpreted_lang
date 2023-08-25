package interpreter

import (
	"fmt"
	"math"
	"strings"

	"github.com/pkg/errors"

	"arc/ast"
)

var (
	UnsupportedOperationError = fmt.Errorf("Invalid operation, unsupported operation, lhs=%s op=%s rhs=%s", "%s", "%s", "%s")
)

type ExpressionInstance struct {
}

var Expression = &ExpressionInstance{}

func (self *ExpressionInstance) Equal(lhs, rhs *ast.RuntimeValue) (bool, error) {
	var r bool
	var rv *ast.RuntimeValue
	var err error

	if lhs.IsOptionKind() {
		lhs = lhs.Unwrap()
	}

	switch l := lhs.Value.(type) {

	case int:
		if rv, err = TypeCoercion.Cast(rhs, ast.IntType); rv != nil {
			r = l == ast.RuntimeValueAs[int](rv)
		}
	case float64:
		if rv, err = TypeCoercion.Cast(rhs, ast.FloatType); rv != nil {
			r = l == ast.RuntimeValueAs[float64](rv)
		}
	case bool:
		if rv, err = TypeCoercion.Cast(rhs, ast.BoolType); rv != nil {
			r = l == ast.RuntimeValueAs[bool](rv)
		}
	case string:
		if !rhs.HasValue() && rhs.IsNoneKind() {
			if !lhs.IsOptionKind(true) {
				return false, errors.New("Invalid operation, rhs is none, and lhs is not option type")
			}
			return false, nil
		}

		if rv, err = TypeCoercion.Cast(rhs, ast.StringType); rv != nil {
			r = l == ast.RuntimeValueAs[string](rv)
		}

		if err != nil {
			return false, err
		}

	case nil:
		if lhs == nil {
			return false, errors.New("Invalid operation, lhs is nil")
		}
		if !lhs.IsOptionKind(true) {
			return false, errors.New("Invalid operation, lhs value is nil, and not an option type")
		}

		if rhs.IsNoneKind() {
			return true, nil
		}

		r = false
	}

	return r, err
}
func (self *ExpressionInstance) NotEqual(lhs, rhs *ast.RuntimeValue) (bool, error) {
	v, err := self.Equal(lhs, rhs)
	if err != nil {
		return false, err
	}
	return !v, nil
}
func (self *ExpressionInstance) GreaterThan(lhs, rhs *ast.RuntimeValue) (bool, error) {
	var r bool
	var rv *ast.RuntimeValue
	var err error

	switch l := lhs.Value.(type) {

	case int:
		if rv, err = TypeCoercion.Cast(rhs, ast.IntType); rv != nil {
			r = l > ast.RuntimeValueAs[int](rv)
		}
	case float64:
		if rv, err = TypeCoercion.Cast(rhs, ast.FloatType); rv != nil {
			r = l > ast.RuntimeValueAs[float64](rv)
		}
	case bool:
		if rv, err = TypeCoercion.Cast(rhs, ast.BoolType); rv != nil {
			r = boolToInt(l) > boolToInt(ast.RuntimeValueAs[bool](rv))
		}
	}

	return r, err
}
func (self *ExpressionInstance) GreaterThanOrEqual(lhs, rhs *ast.RuntimeValue) (bool, error) {
	v, err := self.GreaterThan(lhs, rhs)
	if err != nil {
		return false, err
	}
	if v {
		return v, nil
	}

	v, err = self.Equal(lhs, rhs)
	if err != nil {
		return false, err
	}

	return v, nil
}
func (self *ExpressionInstance) LessThan(lhs, rhs *ast.RuntimeValue) (bool, error) {
	v, err := self.GreaterThan(rhs, lhs)
	if err != nil {
		return false, err
	}

	return !v, nil
}
func (self *ExpressionInstance) LessThanOrEqual(lhs, rhs *ast.RuntimeValue) (bool, error) {
	v, err := self.GreaterThanOrEqual(rhs, lhs)
	if err != nil {
		return false, err
	}

	return !v, nil
}

func (self *ExpressionInstance) And(lhs, rhs *ast.RuntimeValue) (bool, error) {
	var r bool
	var rv *ast.RuntimeValue
	var err error

	switch l := lhs.Value.(type) {

	case int:
		if rv, err = TypeCoercion.Cast(rhs, ast.IntType); rv != nil {
			r = l == ast.RuntimeValueAs[int](rv)
		}
	case float64:
		if rv, err = TypeCoercion.Cast(rhs, ast.FloatType); rv != nil {
			r = l == ast.RuntimeValueAs[float64](rv)
		}
	case bool:
		if rv, err = TypeCoercion.Cast(rhs, ast.BoolType); rv != nil {
			r = l == ast.RuntimeValueAs[bool](rv)
		}
	case string:
		if rv, err = TypeCoercion.Cast(rhs, ast.StringType); rv != nil {
			r = l == ast.RuntimeValueAs[string](rv)
		}
	}

	return r, err
}

func (self *ExpressionInstance) Add(lhs, rhs *ast.RuntimeValue) (*ast.RuntimeValue, error) {
	var res any
	var rv *ast.RuntimeValue
	var err error

	switch r := rhs.Value.(type) {

	case int:
		if rv, err = TypeCoercion.Cast(lhs, ast.IntType); rv != nil {
			res = ast.RuntimeValueAs[int](rv) + r
		}
	case float64:
		if rv, err = TypeCoercion.Cast(lhs, ast.FloatType); rv != nil {
			res = ast.RuntimeValueAs[float64](rv) + r
		}
	case bool:
		if rv, err = TypeCoercion.Cast(lhs, ast.BoolType); rv != nil {
			res = boolToInt(ast.RuntimeValueAs[bool](rv)) + boolToInt(r)
		}
	case string:
		if rv, err = TypeCoercion.Cast(lhs, ast.StringType); rv != nil {
			res = ast.RuntimeValueAs[string](rv) + r
		}
	default:
		return nil, errors.Wrapf(UnsupportedOperationError, lhs.TypeName, "+", rhs.TypeName)
	}
	if err != nil {
		return nil, err
	}
	return ast.NewRuntimeLiteral(res), nil
}
func (self *ExpressionInstance) Sub(lhs, rhs *ast.RuntimeValue) (*ast.RuntimeValue, error) {
	var res any
	var rv *ast.RuntimeValue
	var err error

	switch r := rhs.Value.(type) {

	case int:
		if rv, err = TypeCoercion.Cast(lhs, ast.IntType); rv != nil {
			res = ast.RuntimeValueAs[int](rv) - r
		}
	case float64:
		if rv, err = TypeCoercion.Cast(lhs, ast.FloatType); rv != nil {
			res = ast.RuntimeValueAs[float64](rv) - r
		}
	case bool:
		if rv, err = TypeCoercion.Cast(lhs, ast.BoolType); rv != nil {
			res = boolToInt(ast.RuntimeValueAs[bool](rv)) - boolToInt(r)
		}
	// case string:
	// 	if rv, err = TypeCoercion.Cast(lhs, ast.StringType); rv != nil {
	// 		r =  ast.RuntimeValueAs[string](rv) - r
	// 	}
	default:
		return nil, errors.Wrapf(UnsupportedOperationError, lhs.TypeName, "+", rhs.TypeName)
	}
	if err != nil {
		return nil, err
	}
	return ast.NewRuntimeLiteral(res), nil
}
func (self *ExpressionInstance) Div(lhs, rhs *ast.RuntimeValue) (*ast.RuntimeValue, error) {
	var res any
	var rtv *ast.RuntimeValue
	var err error

	switch r := rhs.Value.(type) {

	case int:
		if rtv, err = TypeCoercion.Cast(lhs, ast.IntType); rtv != nil {
			res = ast.RuntimeValueAs[int](rtv) / r
		}
	case float64:
		if rtv, err = TypeCoercion.Cast(lhs, ast.FloatType); rtv != nil {
			res = ast.RuntimeValueAs[float64](rtv) / r
		}

		// rtv, err := TypeCoercion.Cast(lhs, ast.FloatType)
		// if err != nil {
		// 	return nil, err
		// }
		// if rtv, err = TypeCoercion.Cast(lhs, ast.FloatType); rtv != nil {
		// 	r = ast.RuntimeValueAs[float64](lv) / ast.RuntimeValueAs[float64](rtv)
		// }
	case bool:
		if rtv, err = TypeCoercion.Cast(lhs, ast.BoolType); rtv != nil {
			res = boolToInt(ast.RuntimeValueAs[bool](rtv)) / boolToInt(r)
		}
		// TODO: Maybe we could have some funky string division...
		// TODO: But requires some thought...
		// case string:
		// "abcd" / 2 = "ab"
		// "abcd" / "ab" = ?
		// "abcd" / "ab"(2) = "ab"
		// "abcd" / "ab"(2) = "cd"
	default:
		return nil, errors.Wrapf(UnsupportedOperationError, lhs.TypeName, "/", rhs.TypeName)
	}
	if err != nil {
		return nil, err
	}
	return ast.NewRuntimeLiteral(res), nil
}
func (self *ExpressionInstance) Mul(lhs, rhs *ast.RuntimeValue) (*ast.RuntimeValue, error) {
	var res any
	var rv *ast.RuntimeValue
	var err error

	switch l := lhs.Value.(type) {
	case string:
		// "ab" * 2 = "abab"

		if !rhs.IsNumeric() {
			return nil, &ConstraintError{
				LhsType:        rv.TypeName,
				RhsType:        ast.IntType.TypeName(),
				ConstraintInfo: `"ab" * 2 = "abab" - Lhs is duplicated n times`,
			}
		}

		if rv, err = TypeCoercion.Cast(rhs, ast.IntType); rv != nil {
			res = strings.Repeat(l, rv.Value.(int))
		}

		return ast.NewRuntimeLiteral(res), nil
	}

	switch r := rhs.Value.(type) {

	case int:
		if rv, err = TypeCoercion.Cast(lhs, ast.IntType); rv != nil {
			res = ast.RuntimeValueAs[int](rv) * r
		}
	case float64:
		if rv, err = TypeCoercion.Cast(lhs, ast.FloatType); rv != nil {
			res = ast.RuntimeValueAs[float64](rv) * r
		}
	case bool:
		if rv, err = TypeCoercion.Cast(lhs, ast.BoolType); rv != nil {
			res = boolToInt(ast.RuntimeValueAs[bool](rv)) * boolToInt(r)
		}

	default:
		return nil, errors.Wrapf(UnsupportedOperationError, lhs.TypeName, "*", rhs.TypeName)
	}
	if err != nil {
		return nil, err
	}
	return ast.NewRuntimeLiteral(res), nil
}
func (self *ExpressionInstance) Mod(lhs, rhs *ast.RuntimeValue) (*ast.RuntimeValue, error) {
	var res any
	var rv *ast.RuntimeValue
	var err error

	switch r := rhs.Value.(type) {

	case int:
		if rv, err = TypeCoercion.Cast(lhs, ast.IntType); rv != nil {
			res = ast.RuntimeValueAs[int](rv) % r
		}
	case float64:
		if rv, err = TypeCoercion.Cast(lhs, ast.IntType); rv != nil {
			res = math.Mod(ast.RuntimeValueAs[float64](rv), r)
		}
	default:
		return nil, errors.Wrapf(UnsupportedOperationError, lhs.TypeName, "%", rhs.TypeName)
	}
	if err != nil {
		return nil, err
	}
	return ast.NewRuntimeLiteral(res), nil
}
func (self *ExpressionInstance) Pow(lhs, rhs *ast.RuntimeValue) (*ast.RuntimeValue, error) {
	var res any
	var rv *ast.RuntimeValue
	var err error

	switch lhs.Value.(type) {
	case int, float64:
		var lv *ast.RuntimeValue
		lv, err = TypeCoercion.Cast(lhs, ast.FloatType)
		if err != nil {
			return nil, err
		}
		rv, err = TypeCoercion.Cast(rhs, ast.FloatType)
		if err != nil {
			return nil, err
		}
		res = math.Pow(ast.RuntimeValueAs[float64](lv), ast.RuntimeValueAs[float64](rv))
	default:
		return nil, errors.Wrapf(UnsupportedOperationError, lhs.TypeName, "^", rhs.TypeName)
	}

	if err != nil {
		return nil, err
	}
	return ast.NewRuntimeLiteral(res), nil
}
