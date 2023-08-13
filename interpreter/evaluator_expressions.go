package interpreter

import (
	"github.com/charmbracelet/log"

	"interpreted_lang/ast"
	"interpreted_lang/ast/operators"
)

func getGoType(val any) ast.LiteralKind {
	if val == nil {
		return ast.LiteralKindNull
	}
	switch val.(type) {
	case int:
		return ast.LiteralKindInteger
	case float64:
		return ast.LiteralKindFloat
	case string:
		return ast.LiteralKindString
	case bool:
		return ast.LiteralKindBoolean
	case *ast.Literal:
		return val.(*ast.Literal).Kind
	default:
		log.Fatalf("Unknown literal type: %T", val)
	}

	return ast.LiteralKindUnknown
}

func isType(val interface{}, typeOf interface{}) bool {
	switch typeOf.(type) {
	case int:
		_, ok := val.(int)
		return ok
	case float64:
		_, ok := val.(float64)
		return ok
	case *ast.Literal:
		_, ok := val.(*ast.Literal)
		return ok
	default:
		return false
	}
}

func evaluateIntOperation(
	kind ast.BinaryExpressionKind,
	op operators.Operator,
	l, r int,
) *Result {
	switch kind {
	case ast.BinaryExpressionKindEquality, ast.BinaryExpressionKindRelational:
		return NewResult(operators.BinaryIntComparisonOperation(op, l, r))
	default:
		return NewResult(operators.BinaryIntOperation(op, l, r))
	}
}

func evaluateFloatOperation(
	kind ast.BinaryExpressionKind,
	op operators.Operator,
	l, r float64,
) *Result {
	switch kind {
	case ast.BinaryExpressionKindEquality, ast.BinaryExpressionKindRelational:
		return NewResult(operators.BinaryFloatComparisonOperation(op, l, r))
	default:
		return NewResult(operators.BinaryFloatOperation(op, l, r))
	}
}

func (self *Evaluator) evalCallExpression(node *ast.CallExpression) *Result {
	eval := self.CreateChild()

	var receiver *ast.RuntimeValue
	var fn *ast.FunctionDeclaration

	if node.Receiver != nil {
		if receiver = eval.MustEval(node.Receiver).(*ast.RuntimeValue); receiver != nil {
			fn = receiver.GetMethod(node.FunctionName)
			// fn = self.Env.LookupFunction(node.FunctionName)
			// if fn == nil {
			// 	panic("Undefined function: " + node.FunctionName)
			// }
		} else {
			log.Fatalf("Function receiver is nil: %v", node.Receiver)
		}
	} else {
		fn = self.Env.LookupFunction(node.FunctionName)
		if fn == nil {
			log.Fatalf("Undefined function: %v", node.FunctionName)
		}
	}

	if node.Receiver != nil && fn.Receiver != nil && fn.CustomFuncCb == nil {
		rv := eval.MustEval(node.Receiver).(*ast.RuntimeValue)
		if rv == nil {
			log.Fatalf("Runtime value: %v is nil", rv)
		}
		if rv.Decl == nil {
			log.Fatalf("Runtime value: %v does not have an associated declaration", rv)
		}
		if rv.Kind != ast.RuntimeValueKindObject {
			log.Fatalf("Runtime value: %v is not an object", rv)
		}

		decl, ok := rv.Decl.(*ast.ObjectDeclaration)
		if !ok {
			log.Fatalf("Runtime value: %v is not an object", rv)
		}

		fn, exists := decl.Methods[node.FunctionName]
		if !exists {
			panic("Undefined method: " + node.FunctionName)
		}

		eval.Env.SetVar(fn.Receiver.Name, rv)

	}

	fnArgs := make([]any, len(node.Args))
	if fn.Args != nil {
		for i, param := range fn.Args {
			arg := eval.MustEval(node.Args[i])

			fnArgs[i] = arg
			eval.Env.SetVar(param.Name, arg)
		}
	}

	if fn.CustomFuncCb != nil {
		if fn.Args == nil && node.Args != nil {
			fnArgs = make([]any, len(node.Args))
			for i, arg := range node.Args {
				fnArgs[i] = eval.MustEvalValue(arg)
			}

		}
		return NewResult(fn.CustomFuncCb(fnArgs...))
	}

	result := NewResult()

	if fn.Body != nil {
		result.Merge(eval.Eval(fn.Body))
	}

	return result
}

func (self *Evaluator) evalRangeExpression(node *ast.RangeExpression) *Result {
	r := NewResult()

	lhs := self.MustEval(node.Left)
	rhs := self.MustEval(node.Right)

	self.Env.SetVar("rangeLower", lhs)
	self.Env.SetVar("rangeUpper", rhs)

	return r
}

func evalBinaryOperation(
	kind ast.BinaryExpressionKind,
	op operators.Operator,
	left, right any,
) *Result {
	if left == nil || right == nil {
		return nil
	}

	if isType(left, &ast.Literal{}) {
		left = left.(*ast.Literal).Value
	}
	if isType(right, &ast.Literal{}) {
		right = right.(*ast.Literal).Value
	}

	switch l := left.(type) {
	case int:
		if r, ok := right.(int); ok {
			return evaluateIntOperation(kind, op, l, r)
		}
	case float64:
		if r, ok := right.(float64); ok {
			return evaluateFloatOperation(kind, op, l, r)
		}

	default:
		log.Fatalf("Unsupported operation: %v %s %v", left, op, right)
	}

	return nil
}

func (self *Evaluator) evalBinaryExpression(node *ast.BinaryExpression) *Result {
	left := self.MustEval(node.Left)
	right := self.MustEval(node.Right)

	result := evalBinaryOperation(node.Kind, node.Op, left, right)
	if result != nil {
		return result
	}

	// If left or right have different types, we can't do anything, for ex, (int(1) + float(2.0))
	// But we need to check for this, so we can specifically error out

	// Check mixed type operations
	if (isType(left, int(0)) && isType(right, float64(0))) || (isType(left, float64(0)) && isType(right, int(0))) {
		log.Fatalf("Unsupported operation, arithmetic operations between int and float are not supported, please cast both sides: %v - %s", node, node.AstNode.Token.GetText())
	}

	log.Fatalf("Error evaluating binary expression: %v - %s", self, node.AstNode.Token.GetText())
	panic("Unsupported operation: " + string(node.Op))
}

func (self *Evaluator) evalUnaryExpression(node *ast.UnaryExpression) *Result {
	r := NewResult()

	return r
}

func (self *Evaluator) evalPostfixExpression(node *ast.PostfixExpression) *Result {
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

	log.Fatalf("Error evaluating postfix expression: %v", node)
	panic("Unsupported operation: " + string(node.Op))
}

func (self *Evaluator) evalFieldAccessExpression(node *ast.FieldAccessExpression) *Result {
	r := NewResult()

	if node.StructInstance == nil {
		panic("Struct instance not found")
	}

	switch instance := node.StructInstance.(type) {

	case *ast.Identifier, *ast.VarReference:
		baseObjEval := self.Eval(instance)
		if baseObjEval == nil {
			self.Eval(instance)
			log.Fatalf("Error evaluating field access expression: %v", instance)
		}
		baseObj := baseObjEval.First().(ast.ObjectFieldGetter)
		if baseObj == nil {
			log.Fatalf("Error evaluating field access expression: %v", instance)
		}
		return NewResult(baseObj.GetField(node.FieldName))

	case *ast.FieldAccessExpression:
		resolvedObj := self.Eval(instance)
		if resolvedObj == nil {
			log.Fatalf("Error evaluating field access expression: %v", instance)
		}
		switch ro := resolvedObj.First().(type) {
		case ast.ObjectFieldGetter:
			return NewResult(ro.GetField(node.FieldName))
		case map[string]any:
			return NewResult(ro[node.FieldName])
		}

	default:
		log.Fatalf("Error evaluating field access expression: %v", instance)
	}

	return r
}

func (self *Evaluator) evalAssignmentExpression(node *ast.AssignmentExpression) *Result {
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
			case *ast.ArrayAccessExpression:
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
}

func (self *Evaluator) evalArrayAccessExpression(node *ast.ArrayAccessExpression) *Result {

	r := NewResult()

	valExpr := self.MustEval(node.Instance)
	if valExpr == nil {
		log.Fatalf("Error evaluating array access expression: %v", node)
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
					log.Fatalf("Error evaluating array access expression, start index is greater than end index: %v", node)
				}

				if startIdx < 0 || endIdx > len(arr) {
					log.Fatalf("Error evaluating array access expression, index out of bounds: %v", node)
				}

				r.Add(arr[startIdx:endIdx])
			} else {
				startIdx := startIndex.Value.(int)
				if startIdx < 0 || startIdx >= len(arr) {
					log.Fatalf("Error evaluating array access expression, index out of bounds: %v", node)
				}

				item := arr[startIdx]
				r.Add(item)
			}
		}

	case ast.RuntimeValueKindDict:
		{
			index := self.MustEval(node.StartIndex).(*ast.RuntimeValue)
			if index.Kind != ast.RuntimeValueKindString {
				log.Fatalf("Error evaluating array access expression, index is not a string: %v", node)
			}

			fieldValue := val.GetField(index.Value.(string))

			r.Add(fieldValue)
		}

	default:
		log.Fatalf("Error evaluating array access expression, value is not an array or dict: %v", node)
	}

	return r
}
