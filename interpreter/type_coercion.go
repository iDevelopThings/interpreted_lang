package interpreter

import (
	"strconv"

	"github.com/pkg/errors"

	"arc/ast"
)

type TypeCoercionInstance struct {
}

var TypeCoercion = &TypeCoercionInstance{}

var (
	InvalidCastError      = errors.New("Invalid cast, cannot cast value of type %s to type %s")
	InvalidCastConstraint = errors.New("Invalid constraint, cannot use value type %s here, only %s is possible due to %s")
)

func (t *TypeCoercionInstance) Coerce(lhsRv, rhsRv *ast.RuntimeValue, to *ast.BasicType) (*ast.RuntimeValue, *ast.RuntimeValue) {
	lhs := t.MustCast(lhsRv, to)
	rhs := t.MustCast(rhsRv, to)

	// TODO: Maybe go type sizes should be considered here?
	// For ex, we're always using `int` and `float64` for int and float types
	// Maybe we could use smaller/larger types depending on the values

	return lhs, rhs
}

func (t *TypeCoercionInstance) Cast(rv *ast.RuntimeValue, to *ast.BasicType) (*ast.RuntimeValue, error) {
	if rv.TypeName == to.TypeName() {
		// TODO: Maybe we should actually check here that the underlying value is actually of the correct go type
		return rv, nil
	}

	var newValue any

	switch rv.Kind {

	case ast.RuntimeValueKindInteger:
		switch to {
		case ast.FloatType:
			newValue = float64(rv.Value.(int))
		case ast.StringType:
			newValue = strconv.Itoa(rv.Value.(int))
		case ast.BoolType:
			newValue = rv.Value.(int) != 0
		default:
			return nil, &CastError{rv.TypeName, to.TypeName()}
		}

	case ast.RuntimeValueKindFloat:
		switch to {
		case ast.IntType:
			newValue = int(rv.Value.(float64))
		case ast.StringType:
			newValue = strconv.FormatFloat(rv.Value.(float64), 'f', -1, 64)
		case ast.BoolType:
			newValue = rv.Value.(float64) != 0
		default:
			return nil, &CastError{rv.TypeName, to.TypeName()}
		}

	case ast.RuntimeValueKindBoolean:
		switch to {
		case ast.IntType:
			if rv.Value.(bool) {
				newValue = 1
			} else {
				newValue = 0
			}
		case ast.FloatType:
			if rv.Value.(bool) {
				newValue = 1.0
			} else {
				newValue = 0.0
			}
		case ast.StringType:
			if rv.Value.(bool) {
				newValue = "true"
			} else {
				newValue = "false"
			}
		default:
			return nil, &CastError{rv.TypeName, to.TypeName()}
		}

	default:
		return nil, &CastError{rv.TypeName, to.TypeName()}
	}

	newRv := ast.NewRuntimeValueClone(rv)
	newRv.Kind = ast.RuntimeValueKind(to.Name)
	newRv.TypeName = to.Name
	newRv.Value = newValue

	return newRv, nil
}
func (t *TypeCoercionInstance) CastWithConstraint(lv, rv *ast.RuntimeValue, to *ast.BasicType) (*ast.RuntimeValue, error) {
	if rv.TypeName == to.TypeName() {
		// TODO: Maybe we should actually check here that the underlying value is actually of the correct go type
		return rv, nil
	}

	var newValue any

	switch rv.Kind {

	case ast.RuntimeValueKindInteger:
		switch to {
		case ast.FloatType:
			newValue = float64(rv.Value.(int))
		case ast.StringType:
			newValue = strconv.Itoa(rv.Value.(int))
		case ast.BoolType:
			newValue = rv.Value.(int) != 0
		default:
			return nil, &CastError{rv.TypeName, to.TypeName()}
		}

	case ast.RuntimeValueKindFloat:
		switch to {
		case ast.IntType:
			newValue = int(rv.Value.(float64))
		case ast.StringType:
			newValue = strconv.FormatFloat(rv.Value.(float64), 'f', -1, 64)
		case ast.BoolType:
			newValue = rv.Value.(float64) != 0
		default:
			return nil, &CastError{rv.TypeName, to.TypeName()}
		}

	case ast.RuntimeValueKindBoolean:
		switch to {
		case ast.IntType:
			if rv.Value.(bool) {
				newValue = 1
			} else {
				newValue = 0
			}
		case ast.FloatType:
			if rv.Value.(bool) {
				newValue = 1.0
			} else {
				newValue = 0.0
			}
		case ast.StringType:
			if rv.Value.(bool) {
				newValue = "true"
			} else {
				newValue = "false"
			}
		default:
			return nil, &CastError{rv.TypeName, to.TypeName()}
		}

	// case ast.RuntimeValueKindNone:
	// 	switch to {
	// 	case ast.StringType:
	// 		if lv.IsOptionKind() && !lv.HasValue() {
	//
	// 		}
	// 	}

	default:
		return nil, &CastError{rv.TypeName, to.TypeName()}
	}

	newRv := ast.NewRuntimeValueClone(rv)
	newRv.Kind = ast.RuntimeValueKind(to.Name)
	newRv.TypeName = to.Name
	newRv.Value = newValue

	return newRv, nil
}

func (t *TypeCoercionInstance) MustCast(rv *ast.RuntimeValue, to *ast.BasicType) *ast.RuntimeValue {
	newRv, err := t.Cast(rv, to)
	if err != nil {
		NewErrorAtNode(rv.OriginalNode, err.Error())
	}
	return newRv
}

func (t *TypeCoercionInstance) BothAreRuntimeKind(left *ast.RuntimeValue, right *ast.RuntimeValue, vk ...ast.RuntimeValueKind) bool {
	if left.Kind == right.Kind {
		return true
	}

	for _, k := range vk {
		if left.Kind == k && right.Kind == k {
			return true
		}
	}

	return false
}

func (t *TypeCoercionInstance) BothAreEither(left *ast.RuntimeValue, right *ast.RuntimeValue, vk ...ast.RuntimeValueKind) bool {
	l := left.Kind
	r := right.Kind

	// Both l & r need to be any of the vk, anything else is incorrect

	lis := false
	ris := false

	for _, k := range vk {
		if l == k {
			lis = true
		}
		if r == k {
			ris = true
		}

		if lis && ris {
			return true
		}
	}

	return false
}
