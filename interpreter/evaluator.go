package interpreter

import (
	"github.com/charmbracelet/log"

	"arc/ast"
)

type Evaluator struct {
	Parent *Evaluator

	Env      *Environment
	Children []*Evaluator
}

func NewEvaluator(env *Environment) *Evaluator {
	return &Evaluator{
		Env: env,
	}
}

func (self *Evaluator) GetRoot() *Evaluator {
	if self.Parent == nil {
		return self
	}
	return self.Parent.GetRoot()
}

func (self *Evaluator) CreateChild() *Evaluator {
	rootEnv := self.GetRoot()

	child := &Evaluator{
		Parent: self,
		Env:    self.Env.NewChild(),
	}

	rootEnv.Children = append(rootEnv.Children, child)

	return child
}

func (self *Evaluator) Eval(n any) *Result {
	switch node := n.(type) {
	case *ast.Program:
		return self.evalProgram(node)

	case *ast.ObjectDeclaration:
		return self.evalObjectDeclaration(node)
	case *ast.FunctionDeclaration:
		return self.evalFunctionDeclaration(node)

	case *ast.HttpRouteDeclaration:
		return self.evalHttpRouteDeclaration(node)
	case *ast.HttpServerConfig:
		return self.evalHttpServerConfig(node)
	case *ast.HttpResponseData:
		return self.evalHttpResponseData(node)
	case *ast.HttpRouteBodyInjection:
		return self.evalHttpRouteBodyInjection(node)

	case *ast.IfStatement:
		return self.evalIfStatement(node)
	case *ast.LoopStatement:
		return self.evalLoopStatement(node)
	case *ast.ReturnStatement:
		return self.evalReturnStatement(node)
	case *ast.BreakStatement:
		return self.evalBreakStatement(node)
	case *ast.DeleteStatement:
		return self.evalDeleteStatement(node)
	case *ast.AssignmentStatement:
		return self.evalAssignmentStatement(node)

	// TODO: ContinueStatement
	// case *ast.ContinueStatement:
	// 	return self.evalContinueStatement(node)

	case *ast.PostfixExpression:
		return self.evalPostfixExpression(node)
	case *ast.FieldAccessExpression:
		return self.evalFieldAccessExpression(node)
	case *ast.CallExpression:
		return self.evalCallExpression(node)
	case *ast.BinaryExpression:
		return self.evalBinaryExpression(node)
	case *ast.AssignmentExpression:
		return self.evalAssignmentExpression(node)
	case *ast.UnaryExpression:
		return self.evalUnaryExpression(node)
	case *ast.RangeExpression:
		return self.evalRangeExpression(node)

	case *ast.ObjectInstantiation:
		return self.evalObjectInstantiation(node)
	case *ast.IndexAccessExpression:
		return self.evalArrayAccessExpression(node)
	case *ast.ArrayInstantiation:
		return self.evalArrayInstantiation(node)
	case *ast.DictionaryInstantiation:
		return self.evalDictionaryInstantiation(node)

	case *ast.Literal:
		return self.evalLiteral(node)
	case *ast.Identifier:
		return self.evalIdentifier(node)
	case *ast.VarReference:
		return self.evalVarReference(node)
	case *ast.Block:
		return self.evalBlock(node)
	default:
		if node, ok := n.(*ast.AstNode); ok {
			log.Fatalf("Unhandled AST Node Type: %T - Content: %s", node, node.GetToken())
			return nil
		}
		log.Fatalf("Unhandled AST Node Type: %T", node)
		return nil
	}
}

func (self *Evaluator) MustEval(n any) any {
	res := self.Eval(n)
	if res.HasValue() {
		return res.First()
	}

	log.Fatalf("Error evaluating node: %v", n)

	return nil
}

func (self *Evaluator) MustEvalValue(n any) any {
	value := self.MustEval(n)

	if _, ok := value.(ast.Node); ok {
		switch v := value.(type) {
		case *ast.Literal:
			return v.Value
		default:
			log.Fatalf("Error evaluating node: %v", n)
		}
	} else if rv, ok := value.(*ast.RuntimeValue); ok {
		return rv
	} else if rv, ok := value.([]*ast.RuntimeValue); ok {
		return rv
	} else if result, ok := value.(*Result); ok {
		return result.First()
	} else {
		log.Fatalf("Error evaluating node: %v", n)
	}

	return nil
}

func (self *Evaluator) evalProgram(node *ast.Program) *Result {
	r := NewResult()

	for _, stmt := range node.Statements {
		r.Merge(self.Eval(stmt))
	}

	return r
}

func (self *Evaluator) evalObjectDeclaration(node *ast.ObjectDeclaration) *Result {
	r := NewResult()

	return r
}

func (self *Evaluator) evalFunctionDeclaration(node *ast.FunctionDeclaration) *Result {
	r := NewResult()

	return r
}

func (self *Evaluator) ExecuteFunction(node *ast.FunctionDeclaration) *Result {
	r := NewResult()

	eval := self.CreateChild()

	if node.CustomFuncCb != nil {
		return r.Add(node.CustomFuncCb(eval.Env))
	}

	if node.Args != nil {
		for _, param := range node.Args {
			paramValue := eval.Eval(param)
			if !paramValue.HasValue() {
				log.Fatalf("Function(%s) parameter %s has no value", node.Name, param.Name)
			}
			eval.Env.SetVar(param.Name, paramValue.First())
		}
	}

	if node.Receiver != nil {
		receiver := eval.Eval(node.Receiver)
		if !receiver.HasValue() {
			log.Fatalf("Function(%s) receiver %s has no value", node.Name, node.Receiver.Name)
		}
		eval.Env.SetVar(node.Receiver.Name, receiver.First())
	}

	if node.Body != nil {
		r.Merge(eval.Eval(node.Body))
	}

	return r
}

func (self *Evaluator) evalDictionaryInstantiation(node *ast.DictionaryInstantiation) *Result {
	r := NewResult()

	value := ast.NewRuntimeDictionary()

	for fieldName, expr := range node.Fields {
		fieldVal := self.MustEvalValue(expr)
		value.SetField(fieldName, fieldVal.(*ast.RuntimeValue))
	}

	return r.Add(value)
}

func (self *Evaluator) evalObjectInstantiation(node *ast.ObjectInstantiation) *Result {
	inst := ast.NewRuntimeObject(self.Env.LookupObject(node.TypeName.Name))

	fields := map[string]*ast.RuntimeValue{}
	for fieldName, expr := range node.Fields {
		if lit, ok := expr.(*ast.Literal); ok {
			rv := ast.NewRuntimeValueFromLiteral(lit)
			fields[fieldName] = rv
			continue
		}
		val := self.MustEvalValue(expr)
		if rv, ok := val.(*ast.RuntimeValue); ok {
			fields[fieldName] = rv
			continue
		}
		log.Fatalf("Error evaluating object field: %s, value: %v", fieldName, val)
	}

	inst.Value = fields

	return NewResult(inst)

	// instance := &ast.ObjectRuntimeValue{
	// 	TypeName: node.TypeName,
	// 	Fields:   map[string]any{},
	// 	Decl:     self.Env.LookupObject(node.TypeName),
	// }
	// for fieldName, expr := range node.Fields {
	// 	instance.Fields[fieldName] = self.MustEvalValue(expr)
	// }
	//
	// return NewResult(instance)
}

func (self *Evaluator) evalArrayInstantiation(node *ast.ArrayInstantiation) *Result {
	if node.Type == nil {
		log.Fatalf("Array does not have a type identifier associated with it: %v", node)
	}

	decl := self.Env.LookupObject(node.Type.TypeReference.Type)
	inst := ast.NewRuntimeArray(decl)

	var values []*ast.RuntimeValue

	for _, expr := range node.Values {
		v := self.MustEval(expr)
		values = append(values, v.(*ast.RuntimeValue))
	}

	inst.Value = values

	return NewResult(inst)
}

func (self *Evaluator) evalIfStatement(node *ast.IfStatement) *Result {
	r := NewResult()

	cond := self.MustEval(node.Condition)
	if cond.(bool) {
		body := self.Eval(node.Body)
		r.Merge(body)

		return r
	} else if node.Else != nil {
		switch elsePart := node.Else.(type) {
		case *ast.Block:
			r.Merge(self.Eval(elsePart))
		case *ast.IfStatement:
			r.Merge(self.Eval(elsePart))
		}
	}

	return r
}

func (self *Evaluator) evalLoopStatement(node *ast.LoopStatement) *Result {
	r := NewResult()

	var lower *ast.Literal
	var upper *ast.Literal

	iteratorValue := ast.NewLiteral(node.GetToken(), 0)

	eval := self.CreateChild()

	var rangeArray *ast.RuntimeValue
	var iteratorElement *ast.RuntimeValue

	if node.Range != nil {
		// If we have our range set, we need to eval lhs and rhs and bind their values
		// to the child env
		rangeResult := eval.MustEvalValue(node.Range)
		if rangeResult != nil {
			if rv, ok := rangeResult.(*ast.RuntimeValue); ok && rv.Kind == ast.RuntimeValueKindArray {
				rangeArray = rv
				eval.Env.SetVar("rangeLower", ast.NewLiteral(nil, 0))
				eval.Env.SetVar("rangeUpper", ast.NewLiteral(nil, len(rv.Value.([]*ast.RuntimeValue))-1))

				iteratorElement = rangeArray.Value.([]*ast.RuntimeValue)[0]
			}
		}
		lower = eval.Env.LookupVar("rangeLower").(*ast.Literal)
		upper = eval.Env.LookupVar("rangeUpper").(*ast.Literal)

		if lower == nil || upper == nil {
			log.Fatalf("Upper or lower bound not set for range loop: " + node.GetToken().String())
		}

	}

	setIteratorBinding := func(isInitial bool) {
		if node.As != nil {
			// Set `as i` to the lower bound
			if isInitial {
				iteratorValue.Value = lower.Value
			}
			if rangeArray != nil {
				eval.Env.SetVar(node.As.Name, iteratorElement)
			} else {
				eval.Env.SetVar(node.As.Name, iteratorValue)
			}
		} else {
			// Automatically bind `it` to the lower bound

			if rangeArray != nil {
				eval.Env.SetVar("it", iteratorElement)
			} else {
				eval.Env.SetVar("it", iteratorValue)
			}
		}
	}

	setIteratorBinding(true)

	for {
		// Now we need to check the lower and upper bounds and
		// see if we need to stop looping

		// Eval the body
		bodyResult := eval.Eval(node.Body)
		if bodyResult != nil {
			r.Merge(bodyResult)
		}

		// Check the break status
		if result, ok := bodyResult.HasStatus(ResultStatusBreak); ok {
			r.AddExisting(result)
			break
		}

		// Check the continue status
		if result, ok := bodyResult.HasStatus(ResultStatusContinue); ok {
			r.AddExisting(result)
			continue
		}

		if node.Range != nil {
			if node.Step != nil {
				iteratorValue.Add(node.Step)
			} else {
				iteratorValue.Add(1)
			}

			if rangeArray != nil {
				vals := rangeArray.Value.([]*ast.RuntimeValue)
				if iteratorValue.Value.(int) > len(vals)-1 {
					break
				}
				iteratorElement = rangeArray.Value.([]*ast.RuntimeValue)[iteratorValue.Value.(int)]
				setIteratorBinding(false)
			}
		}

		if node.Range != nil {
			if iteratorValue.IsGreaterThan(upper) {
				break
			}
		}

	}

	return r
}

func (self *Evaluator) evalAssignmentStatement(node *ast.AssignmentStatement) *Result {
	r := NewResult()

	var value any
	if node.Value != nil {
		value = self.MustEval(node.Value)
	}

	self.Env.SetVar(node.Name.Name, value)

	return r.Add(value)
}

func (self *Evaluator) evalReturnStatement(node *ast.ReturnStatement) *Result {
	r := NewResult()

	var val any

	if node.Value != nil {
		val = self.Eval(node.Value)
	}

	return r.Add(val, ResultStatusReturn)
}

func (self *Evaluator) evalBreakStatement(node *ast.BreakStatement) *Result {
	return NewResultWithStatus(nil, ResultStatusBreak)
}

func (self *Evaluator) evalDeleteStatement(node *ast.DeleteStatement) *Result {
	r := NewResult()

	var parentVal *ast.RuntimeValue
	var element *ast.RuntimeValue

	var startIndex *ast.RuntimeValue
	var endIndex *ast.RuntimeValue
	var elements []*ast.RuntimeValue

	var fieldName string

	switch what := node.What.(type) {

	case *ast.VarReference:
		self.Env.DeleteVar(node.What.(*ast.VarReference).Name)
		return r

	case *ast.IndexAccessExpression:
		parentVal = self.MustEval(what.Instance).(*ast.RuntimeValue)

		if what.IsSlice {
			elements = self.MustEval(what).([]*ast.RuntimeValue)
			startIndex = self.MustEval(what.StartIndex).(*ast.RuntimeValue)
			if what.EndIndex != nil {
				endIndex = self.MustEval(what.EndIndex).(*ast.RuntimeValue)
			} else {
				endIndex = ast.NewRuntimeLiteral(len(elements) - 1)
			}
		} else {
			startIndex = self.MustEval(what.StartIndex).(*ast.RuntimeValue)
			element = startIndex
		}
		// if element.Kind != ast.RuntimeValueKindString {
		// 	log.Fatalf("Cannot delete element with non-string index: " + node.AstNode.Token.GetText())
		// }

	case *ast.FieldAccessExpression:
		parentVal = self.MustEval(what.StructInstance).(*ast.RuntimeValue)

		fieldName = what.FieldName
		element = parentVal.GetField(what.FieldName)
	}

	if parentVal == nil {
		log.Fatalf("Cannot delete nil value - %s", node.GetToken())
	}

	switch parentVal.Kind {

	case ast.RuntimeValueKindArray:
		if elements != nil {
			ok, err := parentVal.RemoveArrayElementsInRange(startIndex.Value.(int), endIndex.Value.(int))
			if !ok || err != nil {
				log.Fatalf("Error deleting array element: " + err.Error())
			}
		} else {
			ok, err := parentVal.RemoveArrayElementByIndex(element)
			if !ok || err != nil {
				if err != nil {
					log.Fatalf("Error deleting array element: " + err.Error())
				}
				log.Fatalf("Cannot delete element with non-string index: " + node.GetToken().String())
			}
		}

	case ast.RuntimeValueKindDict:
		if fieldName == "" {
			parentVal.DeleteDictElement(element.Value.(string))
		} else {
			parentVal.DeleteDictElement(fieldName)
		}

	}

	// what := self.MustEvalValue(node.What)
	// if what == nil {
	// 	log.Fatalf("Cannot delete nil value - %s", node.AstNode.Token.GetText())
	// }

	return r
}

func (self *Evaluator) evalLiteral(node *ast.Literal) *Result {
	r := NewResult()

	if node == nil {
		return r
	}
	if err := node.Unescape(); err != nil {
		panic("Error evaluating literal(unescape): " + err.Error())
	}

	rv := ast.NewRuntimeValueFromLiteral(node)

	return r.Add(rv)
}

func (self *Evaluator) evalIdentifier(node *ast.Identifier) *Result {
	r := NewResult()

	val := self.Env.LookupVar(node.Name)

	return r.Add(val)
}

func (self *Evaluator) evalVarReference(node *ast.VarReference) *Result {
	value := self.Env.LookupVar(node.Name)
	if value == nil {
		panic("Undefined variable: " + node.Name)
	}

	return NewResult(value)
}

func (self *Evaluator) evalBlock(node *ast.Block) *Result {
	r := NewResult()

	for _, stmtOrExpr := range node.Statements {
		switch se := stmtOrExpr.(type) {
		case *ast.ReturnStatement, *ast.BreakStatement:
			stmt := self.Eval(se)
			r.Merge(stmt)
		case *ast.IfStatement:
			stmt := self.Eval(se)
			r.Merge(stmt)
		case ast.Statement:
			self.Eval(se)
		case ast.Expr:
			self.Eval(se)
		default:
			panic("Unknown type in function body")
		}

		if r.HasAnyStatus(ResultStatusReturn, ResultStatusBreak, ResultStatusContinue) {
			return r
		}
	}

	return r
}
