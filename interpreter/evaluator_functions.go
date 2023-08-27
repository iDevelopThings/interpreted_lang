package interpreter

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/log"

	"arc/ast"
	"arc/interpreter/errors"
)

func (self *Evaluator) evalDeferStatement(node *ast.DeferStatement) *Result {
	r := NewResult()

	if getFrame() == nil {
		log.Errorf("Cannot defer statement, no stack frame")
		return r
	}

	getFrame().deferStack.Push(node)

	return r
}

func (self *Evaluator) lookupFuncDeclaration(node *ast.CallExpression) (receiver *ast.RuntimeValue, fn *ast.FunctionDeclaration) {

	if node.Receiver != nil {
		// If we're accessing a static method, our receiver is a TypeReference to the type
		// that the method is defined on...
		// So we'll resolve the fn declaration from the type reference

		if node.IsStaticAccess {
			typeRef, ok := node.Receiver.(*ast.TypeReference)
			if !ok {
				NewErrorAtNode(node.Receiver, "Receiver is not a type reference: %#v", node.Receiver)
			}
			receiverObj := Registry.LookupType(typeRef.Type)
			if receiverObj == nil {
				NewErrorAtNode(typeRef, "Undefined object: %#v", typeRef.Type)
			}

			switch obj := receiverObj.(type) {
			case *ast.ObjectDeclaration:
				fn = obj.GetMethod(node.Function.Name)
			case *ast.EnumDeclaration:
				receiver = self.Env.GetRoot().LookupVar(node.Receiver.(*ast.TypeReference).Type).(*ast.RuntimeValue)
				// receiver = eval.MustEval(node.Receiver).(*ast.RuntimeValue)
				if receiver == nil {
					NewErrorAtNode(node, "Enum receiver is nil: %v", node.Receiver)
				}
				fn = receiver.GetMethod(node.Function.Name)
			default:
				NewErrorAtNode(typeRef, "Invalid receiver type: %#v", typeRef.Type)
			}

			return
		}

		// Otherwise, we're accessing an instance method, so we need to resolve the variable or whatever
		// to first find its type and then resolve the method from that type

		receiver = self.MustEval(node.Receiver).(*ast.RuntimeValue)
		if receiver == nil {
			NewErrorAtNode(node, "Receiver is nil: %v", node.Receiver)
		}
		fn = receiver.GetMethod(node.Function.Name)
		if fn == nil {
			NewErrorAtNode(node.Function, "Failed to resolve function: %v", node.Function.Name)
		}

		return
	}

	fn = Registry.LookupFunction(node.Function.Name)
	if fn == nil {
		NewErrorAtNode(node.Function, "Undefined function: %v", node.Function.Name)
	}
	return
}

func (self *Evaluator) evalCallExpression(node *ast.CallExpression) *Result {
	r := NewResult()
	eval := self.CreateChild()

	_, fn := eval.lookupFuncDeclaration(node)

	frame := eval.Env.NewFrame(fn)
	frame.addCallSite(node)
	defer frame.pop()

	eval.bindWrappedFunction(
		fn,
		frame,
		r,
		node,
	)

	return r
}

func (self *Evaluator) bindWrappedFunction(
	fn *ast.FunctionDeclaration,
	frame *StackFrame,
	r *Result,
	node *ast.CallExpression,
) {
	self.executeWrappedFunction(fn, frame, r, node)

	if frame.didError() {
		defer errors.SetStrategy(errors.ExitOnError)

		log.Debugf("Error dump:")

		// ErrorManager.SetNode(node)
		errors.SetStrategy(errors.ContinueOnError)

		for _, err := range frame.errs {
			log.Errorf("Error: %s", err.Error)
		}

		fmt.Printf(strings.Repeat("-", 80) + "\n")
		log.Debugf("Callsites:")

		frames := frame.unrollCallSites()

		frameIdx := 0
		for _, stackFrame := range frames {

			for _, callSite := range stackFrame.callSites {
				NewErrorAtNode(callSite, "(frame:%d)Call site: %d", stackFrame.id, frameIdx)
				frameIdx++
			}

		}

		// NewError("Error executing function \"%s\"\nerror:", fn.Name)
		// NewError("%s", err.Error)
		// NewError("At node:%s\n", err.Node.GetToken().Value)
	}

}

func (self *Evaluator) executeWrappedFunction(
	fn *ast.FunctionDeclaration,
	frame *StackFrame,
	r *Result,
	node *ast.CallExpression,
) {
	defer func() {
		if r := recover(); r != nil {
			frame.errored(r)
		}

		if frame.deferStack.Len() > 0 {
			for frame.deferStack.Len() > 0 {
				deferStmt := frame.deferStack.Pop()
				if deferStmt == nil {
					continue
				}
				self.Eval(deferStmt.Func.Body)
			}
		}

	}()

	// Now... if we're not calling a static method, we need to set the receiver variable
	// for ex: func (x MyType) foo() { ... }
	// We need to set `x` to our struct instance value on the environment scope

	// If our receiver is nil, we're calling a global fn
	if fn.CustomFuncCb == nil && !node.IsStaticAccess && node.Receiver != nil {
		// Evaluating the receiver will resolve our variable to its runtime value
		// Which we'll then apply to the fn scope as the receiver variable
		rv := self.MustEval(node.Receiver).(*ast.RuntimeValue)
		if rv == nil {
			NewErrorAtNode(node.Receiver, "Runtime value: %v is nil", rv)
		}
		if rv.Decl == nil {
			NewErrorAtNode(node.Receiver, "Runtime value: %v does not have an associated declaration", rv)
		}
		if rv.Kind != ast.RuntimeValueKindObject {
			NewErrorAtNode(node.Receiver, "Runtime value: %v is not an object", rv)
		}

		decl, ok := rv.Decl.(*ast.ObjectDeclaration)
		if !ok {
			NewErrorAtNode(node.Receiver, "Runtime value: %v is not an object", rv)
		}

		// Now, even though we've resolved the function above, we'll make
		// sure it actually exists on this variable's runtime object/value

		fnDecl, exists := decl.Methods[node.Function.Name]
		if !exists {
			panic("Undefined method: " + node.Function.Name)
		}
		if fn != fnDecl {
			NewErrorAtNode(node.Function, "Resolved function: %v is not the same as the function declaration: %v", fn, fnDecl)
		}

		self.Env.SetVar(fn.Receiver.Name, rv)
	}

	fnArgs := make([]any, 0)
	if len(fn.Args) > 0 {
		var declArg *ast.TypedIdentifier
		var varArgList []any
		for i, arg := range node.Args {
			if declArg == nil || !declArg.TypeReference.IsVariadic {
				declArg = fn.Args[i]
			}

			argValue := self.MustEval(arg)
			if declArg.TypeReference.IsVariadic {
				fnArgs = append(fnArgs, argValue)
				if self.Env.LookupVar(declArg.Name) == nil {
					self.Env.SetVar(declArg.Name, varArgList)
				}
			} else {
				self.Env.SetVar(declArg.Name, argValue)
				fnArgs = append(fnArgs, argValue)
			}

		}
	}

	if fn.CustomFuncCb != nil {
		if fn.Args == nil && node.Args != nil {
			fnArgs = make([]any, len(node.Args))
			for i, arg := range node.Args {
				fnArgs[i] = self.MustEvalValue(arg)
			}
		}

		fnArgs = append([]any{self.Env}, fnArgs...)
		r.Add(fn.CustomFuncCb(fnArgs...))
	} else {

		if fn.Body != nil {
			result := self.Eval(fn.Body)

			if fn.ReturnType.IsOptionType {
				self.wrapResultInOptionType(result, fn.ReturnType)
			}

			r.Merge(result)
		}

	}

}

// This is only used by the interpreter to execute the main function on boot
func (self *Evaluator) ForceExecuteFunction(node *ast.FunctionDeclaration) *Result {
	r := NewResult()

	eval := self.CreateChild()
	frame := eval.Env.NewFrame(node)
	defer frame.pop()

	eval.bindForceWrappedFunction(node, frame, r)

	return r
}

func (self *Evaluator) bindForceWrappedFunction(
	node *ast.FunctionDeclaration,
	frame *StackFrame,
	r *Result,
) {
	self.executeForceWrappedFunction(node, frame, r)

	if frame.didError() {
		for _, err := range frame.errs {
			log.Errorf("Error dump:")
			log.Errorf("Error executing function \"%s\"\nerror:", node.Name)
			log.Errorf("%s", err.Error)
			log.Errorf("At node:%s\n", err.Node.GetToken().Value)
		}
	}

}

func (self *Evaluator) executeForceWrappedFunction(
	node *ast.FunctionDeclaration,
	frame *StackFrame,
	r *Result,
) {
	defer func() {
		if r := recover(); r != nil {
			frame.errored(r)
		}

		if frame.deferStack.Len() > 0 {
			for frame.deferStack.Len() > 0 {
				deferStmt := frame.deferStack.Pop()
				if deferStmt == nil {
					continue
				}

				self.Eval(deferStmt.Func)
			}
		}

	}()

	if node.CustomFuncCb != nil {
		r.Add(node.CustomFuncCb(self.Env))
	}

	if node.Args != nil {
		for _, param := range node.Args {
			paramValue := self.Eval(param)
			if !paramValue.HasValue() {
				NewErrorAtNode(param, "Function(%s) parameter %s has no value", node.Name, param.Name)
			}
			self.Env.SetVar(param.Name, paramValue.First())
		}
	}

	if node.Receiver != nil {
		receiver := self.Eval(node.Receiver)
		if !receiver.HasValue() {
			NewErrorAtNode(node.Receiver, "Function(%s) receiver %s has no value", node.Name, node.Receiver.Name)
		}
		self.Env.SetVar(node.Receiver.Name, receiver.First())
	}

	if node.Body != nil {
		r.Merge(self.Eval(node.Body))
	}
}
