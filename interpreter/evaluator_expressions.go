package interpreter

import (
	"fmt"

	"arc/ast"
	"arc/ast/operators"
)

func boolToInt(v bool) int {
	if v {
		return 1
	}
	return 0
}
func boolToFloat(v bool) float64 {
	if v {
		return 1.0
	}
	return 0.0
}

func evalBinaryOperation(
	mainNode ast.Node, // Only used for error reporting
	kind ast.BinaryExpressionKind,
	op operators.Operator,
	originalLeft, originalRight *ast.RuntimeValue,
	leftNode, rightNode ast.Expr,
) *Result {

	var rawResult any
	var opError error

	switch kind {

	case ast.BinaryExpressionKindRegular:
		{

			switch op {
			case operators.PlusPlus:
				rawResult, opError = Expression.Add(originalLeft, ast.NewRuntimeLiteral(1))
			case operators.MinusMinus:
				rawResult, opError = Expression.Sub(originalLeft, ast.NewRuntimeLiteral(1))
			case operators.Plus:
				rawResult, opError = Expression.Add(originalLeft, originalRight)
			case operators.Minus:
				rawResult, opError = Expression.Sub(originalLeft, originalRight)
			case operators.Multiply:
				rawResult, opError = Expression.Mul(originalLeft, originalRight)
			case operators.Divide:
				rawResult, opError = Expression.Div(originalLeft, originalRight)
			case operators.Modulo:
				rawResult, opError = Expression.Mod(originalLeft, originalRight)
			case operators.Power:
				rawResult, opError = Expression.Pow(originalLeft, originalRight)
			}

		}

	case ast.BinaryExpressionKindComparison:
		{

			switch op {

			case operators.EqualEqual:
				rawResult, opError = Expression.Equal(originalLeft, originalRight)

			case operators.NotEqual:
				rawResult, opError = Expression.NotEqual(originalLeft, originalRight)

			case operators.GreaterThan:
				rawResult, opError = Expression.GreaterThan(originalLeft, originalRight)

			case operators.GreaterThanOrEqual:
				rawResult, opError = Expression.GreaterThanOrEqual(originalLeft, originalRight)

			case operators.LessThan:
				rawResult, opError = Expression.LessThan(originalLeft, originalRight)

			case operators.LessThanOrEqual:
				rawResult, opError = Expression.LessThanOrEqual(originalLeft, originalRight)

			case operators.And:
				rawResult, opError = Expression.And(originalLeft, originalRight)

			default:
				NewErrorAtNode(mainNode, "Error evaluating binary expression, unhandled comparison operator: %v", op)
			}
		}

	case ast.BinaryExpressionKindAssignment:
		{
			// We should only only be handling the right side of the assignment here
			// The `evalAssignmentStatement` will handle finding the correct var on the left
			// and assigning the value to it

			var er *Result

			switch op {

			case operators.PlusEqual:
				er = evalBinaryOperation(mainNode, ast.BinaryExpressionKindRegular, operators.Plus, originalLeft, originalRight, leftNode, rightNode)
			case operators.MinusEqual:
				er = evalBinaryOperation(mainNode, ast.BinaryExpressionKindRegular, operators.Minus, originalLeft, originalRight, leftNode, rightNode)
			case operators.MultiplyEqual:
				er = evalBinaryOperation(mainNode, ast.BinaryExpressionKindRegular, operators.Multiply, originalLeft, originalRight, leftNode, rightNode)
			case operators.DivideEqual:
				er = evalBinaryOperation(mainNode, ast.BinaryExpressionKindRegular, operators.Divide, originalLeft, originalRight, leftNode, rightNode)
			default:
				NewErrorAtNode(mainNode, "Error evaluating binary expression, unhandled assignment operator: %v", op)
			}

			if er == nil {
				NewErrorAtNode(mainNode, "Error evaluating binary expression, result is nil: %v", op)
			}

			rt := er.First().(*ast.RuntimeValue)
			originalLeft.Apply(rt)

			return NewResult(originalLeft)
		}

	}

	if opError != nil {
		NewErrorAtNode(mainNode, "Error evaluating binary expression: %s", opError.Error())
		return nil
	}

	if rv, ok := rawResult.(*ast.RuntimeValue); ok {
		return NewResult(rv)
	}

	rv := ast.NewRuntimeLiteral(rawResult)

	return NewResult(rv)
}

/*func (self *Evaluator) evalAssignmentExpression(node *ast.AssignmentExpression) *Result {
	r := NewResult()

	varRef := self.MustEvalValue(node.Left)
	if varRef == nil {
		log.Fatalf("Error evaluating postfix expression: %v", node)
	}
	rightRef := self.MustEvalValue(node.Value)

	if _, ok := varRef.(*ast.RuntimeValue); !ok {
		log.Fatalf("Error evaluating postfix expression, varRef is not a runtime value: %v", node)
	}
	left := varRef.(*ast.RuntimeValue)
	if _, ok := rightRef.(*ast.RuntimeValue); !ok {
		log.Fatalf("Error evaluating postfix expression, right is not a runtime value: %v", node)
	}
	right := rightRef.(*ast.RuntimeValue)

	var result *Result

	if node.Op == operators.Equal {
		if left == nil {
			switch accessor := node.Left.(type) {
			case *ast.IndexAccessExpression:
				left = self.MustEvalValue(accessor.Instance).(*ast.RuntimeValue)
				if left != nil && left.Kind == ast.RuntimeValueKindDict {
					dictKey := self.MustEvalValue(accessor.StartIndex).(*ast.RuntimeValue)
					left.SetField(dictKey.Value.(string), right)
					return NewResult(left)
				} else {
					log.Fatalf("Error evaluating Assignment expression, left is nil: %v", node)
				}

			case *ast.FieldAccessExpression:
				left = self.MustEvalValue(accessor.StructInstance).(*ast.RuntimeValue)
				if left != nil && left.Kind == ast.RuntimeValueKindDict {
					left.SetField(accessor.FieldName, right)
					return NewResult(left)
				} else {
					log.Fatalf("Error evaluating Assignment expression, left is nil: %v", node)
				}

			default:
				log.Fatalf("Error evaluating Assignment expression, left is nil: %v", node)

			}

		}

		if left.Kind != right.Kind {
			log.Fatalf("Error evaluating Assignment expression, left and right types are not the same: %v, lhs: %v, rhs: %v", node, left, right)
		}

		result = NewResult(right)
	} else {
		result = evalBinaryOperation(
			ast.BinaryExpressionKindAssignment,
			node.Op,
			varRef,
			right,
		)
	}

	if result != nil {
		resultValue := result.First()
		if rv, ok := resultValue.(*ast.RuntimeValue); ok {
			left.Value = rv.Value
		} else {
			log.Fatalf("Error evaluating Assignment expression, result is not a runtime value: %v", node)
		}

		return NewResult(varRef)
	}

	return r
}*/

func (self *Evaluator) evalBinaryExpression(node *ast.BinaryExpression) *Result {
	left := self.MustEval(node.Left).(*ast.RuntimeValue)
	right := self.MustEval(node.Right).(*ast.RuntimeValue)

	result := evalBinaryOperation(
		node, node.Kind,
		node.Op,
		left, right,
		node.Left, node.Right,
	)
	if result != nil {
		return result
	}

	// If left or right have different types, we can't do anything, for ex, (int(1) + float(2.0))
	// But we need to check for this, so we can specifically error out

	// Check mixed type operations
	// if (isType(left, int(0)) && isType(right, float64(0))) || (isType(left, float64(0)) && isType(right, int(0))) {
	// 	log.Warnf("Unsupported operation, arithmetic operations between int and float are not supported, please cast both sides: %v - %s", node, node.GetToken())
	// }
	NewErrorAtNode(node, "Error evaluating binary expression")

	return nil
}

func (self *Evaluator) evalUnaryExpression(node *ast.UnaryExpression) *Result {
	r := NewResult()

	left := self.MustEval(node.Left).(*ast.RuntimeValue)

	var newValue *ast.RuntimeValue

	switch node.Op {
	case operators.PlusPlus, operators.MinusMinus:
		{
			r := evalBinaryOperation(
				node, ast.BinaryExpressionKindRegular,
				node.Op,
				left, nil,
				node.Left, nil,
			)
			if r == nil || !r.HasValue() {
				NewErrorAtNode(node, "Error evaluating unary expression: %v", node)
			}

			newValue = r.First().(*ast.RuntimeValue)

			left.Apply(newValue)
		}

	default:
		NewErrorAtNode(node, "Unhandled unary expression: %v", node)
	}

	if newValue == nil {
		NewErrorAtNode(node, "Error evaluating unary expression: %v", node)
	}

	return r.Add(left)
}

/*func (self *Evaluator) evalPostfixExpression(node *ast.PostfixExpression) *Result {
	varRef := self.MustEval(node.Left)
	if varRef == nil {
		log.Fatalf("Error evaluating postfix expression: %v", node)
	}

	var left any = varRef
	var newValue any = nil

	if isType(varRef, &ast.Literal{}) {
		left = left.(*ast.Literal).Value
	}

	switch l := left.(type) {
	case int:
		newValue = operators.BinaryIntOperation(node.Op, l, 1)
	case float64:
		newValue = operators.BinaryFloatOperation(node.Op, l, 1)
	}

	if newValue != nil {
		if !isType(varRef, &ast.Literal{}) {
			log.Fatalf("Error evaluating postfix expression, varRef is not a literal type so updated value cannot be stored: %v", node)
		}

		varRef.(*ast.Literal).Value = newValue

		return NewResult(varRef)
	}

	NewErrorAtNode(node, "Error evaluating postfix expression: %v", node)

	// panic("Unsupported operation: " + string(node.Op))
	return nil
}*/

func (self *Evaluator) evalFieldAccessExpression(node *ast.FieldAccessExpression) *Result {
	r := NewResult()

	if node.StructInstance == nil {
		NewErrorAtNode(node, "Struct instance not found")
	}

	switch instance := node.StructInstance.(type) {

	case *ast.Identifier, *ast.VarReference:
		baseObjEval := self.Eval(instance)
		if baseObjEval == nil {
			self.Eval(instance)
			NewErrorAtNode(instance, "Error evaluating field access expression: %v", instance)
		}
		baseObj := baseObjEval.First().(ast.ObjectFieldGetter)
		if baseObj == nil {
			NewErrorAtNode(instance, "Error evaluating field access expression: %v", instance)
		}
		return NewResult(baseObj.GetField(node.FieldName))

	case *ast.FieldAccessExpression:
		resolvedObj := self.Eval(instance)
		if resolvedObj == nil {
			NewErrorAtNode(instance, "Error evaluating field access expression: %v", instance)
		}
		switch ro := resolvedObj.First().(type) {
		case ast.ObjectFieldGetter:
			return NewResult(ro.GetField(node.FieldName))
		case map[string]any:
			return NewResult(ro[node.FieldName])
		}

	case *ast.IndexAccessExpression:
		resolved := self.MustEvalValue(instance)
		if resolved == nil {
			NewErrorAtNode(instance, "Error evaluating field access expression: %v", instance)
		}
		switch ro := resolved.(type) {
		case ast.ObjectFieldGetter:
			return NewResult(ro.GetField(node.FieldName))
		case map[string]any:
			return NewResult(ro[node.FieldName])
		}
		return r.Add(resolved)

	default:
		NewErrorAtNode(instance, "Error evaluating field access expression: %v", instance)
	}

	return r
}

func (self *Evaluator) evalIndexAccessExpression(node *ast.IndexAccessExpression) *Result {

	r := NewResult()

	valExpr := self.MustEval(node.Instance)
	if valExpr == nil {
		NewErrorAtNode(node, "Error evaluating array access expression: %v", node)
	}
	val := valExpr.(*ast.RuntimeValue)

	switch val.Kind {
	case ast.RuntimeValueKindArray:
		{
			arr := val.Value.([]*ast.RuntimeValue)

			startIndex := self.MustEval(node.StartIndex).(*ast.RuntimeValue)

			if node.IsSlice {
				startIdx := startIndex.Value.(int)
				endIdx := len(arr)
				if node.EndIndex != nil {
					endIndex := self.MustEval(node.EndIndex).(*ast.RuntimeValue)
					endIdx = endIndex.Value.(int)
				}

				if startIdx > endIdx {
					NewErrorAtNode(node, "Error evaluating array access expression, start index is greater than end index: %v", node)
				}

				if startIdx < 0 || endIdx > len(arr) {
					NewErrorAtNode(node, "Error evaluating array access expression, index out of bounds: %v", node)
				}

				r.Add(arr[startIdx:endIdx])
			} else {
				startIdx := startIndex.Value.(int)
				if startIdx < 0 || startIdx >= len(arr) {
					NewErrorAtNode(node, "Error evaluating array access expression, index out of bounds: %v", node)
				}

				item := arr[startIdx]
				r.Add(item)
			}
		}

	case ast.RuntimeValueKindDict, ast.RuntimeValueKindEnumValue:
		{
			index := self.MustEval(node.StartIndex).(*ast.RuntimeValue)
			if index.Kind != ast.RuntimeValueKindString && val.Kind == ast.RuntimeValueKindDict {
				NewErrorAtNode(node, "Error evaluating array access expression, index is not a string: %v", node)
			}

			fieldValue := val.GetField(fmt.Sprintf("%v", index.Value))

			r.Add(fieldValue)
		}

	default:
		NewErrorAtNode(node, "Error evaluating array access expression, value is not an array or dict: %v", node)
	}

	return r
}

func (self *Evaluator) evalRangeExpression(node *ast.RangeExpression) *Result {
	r := NewResult()

	lhs := self.MustEval(node.Left)
	rhs := self.MustEval(node.Right)

	self.Env.SetVar("rangeLower", lhs)
	self.Env.SetVar("rangeUpper", rhs)

	return r
}
