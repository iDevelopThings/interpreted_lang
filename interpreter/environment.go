package interpreter

import (
	"arc/ast"
	"arc/log"
)

type FunctionTypeCallback = func(args ...any) any

type Environment struct {
	parent *Environment
	vars   map[string]any
}

var RootEnvironment *Environment

func NewEnvironment() *Environment {
	e := &Environment{
		vars: map[string]any{},
	}

	return e
}

func (self *Environment) NewFrame(function ...*ast.FunctionDeclaration) *StackFrame {
	frame := NewStackFrame(getFrame())
	frame.env = self
	if len(function) > 0 {
		frame.function = function[0]
	}

	return frame
}

func (self *Environment) NewChild() *Environment {
	env := NewEnvironment()
	env.parent = self

	return env
}

func (self *Environment) LookupVar(name string) any {
	if getFrame() != nil {
		if v, ok := getFrame().vars[name]; ok {
			return v
		}
	}

	if v, ok := self.vars[name]; ok {
		return v
	}
	if self.parent != nil {
		return self.parent.LookupVar(name)
	}
	return nil
}

func (self *Environment) SetVar(name string, value any) {
	self.vars[name] = value
	if getFrame() != nil {
		if rt, ok := value.(*ast.RuntimeValue); ok {
			getFrame().vars[name] = rt
		} else {
			log.Warnf("SetVar: trying to push a runtime value to the stack but it's not a runtime value: %s, value=%#v", name, value)
		}
	}
}
func (self *Environment) AppendVars(vars map[string]any) {
	for k, v := range vars {
		self.SetVar(k, v)
	}
}
func (self *Environment) DeleteVar(name string) bool {
	if getFrame() != nil {
		if v, ok := getFrame().vars[name]; ok {
			v.Value = nil
			log.Debugf("Deleted var %s from current frame", name)
			delete(getFrame().vars, name)
			return true
		}
	}

	if v, ok := self.vars[name]; ok {
		if rv, ok := v.(*ast.RuntimeValue); ok {
			rv.Value = nil
		}

		log.Debugf("Deleted var %s from environment", name)

		delete(self.vars, name)

		return true
	}

	return false
}

// func (self *Environment) LookupFunction(name string) *ast.FunctionDeclaration {
// 	// First we'll do a quick check on the root environment
// 	// If it exists(since it could be a global function), we'll return it
// 	// This will reduce the stack size for the recursive lookup
//
// 	if RootEnvironment != nil {
// 		if f, ok := RootEnvironment.functions[name]; ok {
// 			return f
// 		}
// 	}
//
// 	if f, ok := self.functions[name]; ok {
// 		return f
// 	}
// 	if self.parent != nil {
// 		return self.parent.LookupFunction(name)
// 	}
// 	return nil
// }

// func (self *Environment) LookupObject(name string) *ast.ObjectDeclaration {
// 	// First we'll do a quick check on the root environment
// 	// If it exists(since it could be a global object), we'll return it
// 	// This will reduce the stack size for the recursive lookup
//
// 	if RootEnvironment != nil {
// 		if o, ok := RootEnvironment.objects[name]; ok {
// 			return o
// 		}
// 		if bt, ok := ast.BasicTypes[name]; ok {
// 			return &ast.ObjectDeclaration{
// 				AstNode: bt.GetAstNode(),
// 				Name:    ast.NewIdentifierWithValue(nil, bt.TypeName()),
// 			}
// 		}
// 	}
//
// 	if o, ok := self.objects[name]; ok {
// 		return o
// 	}
// 	if self.parent != nil {
// 		return self.parent.LookupObject(name)
// 	}
//
// 	if bt, ok := ast.BasicTypes[name]; ok {
// 		return &ast.ObjectDeclaration{
// 			AstNode: bt.GetAstNode(),
// 			Name:    ast.NewIdentifierWithValue(nil, bt.TypeName()),
// 		}
// 	}
//
// 	return nil
// }

// func (self *Environment) SetFunction(function *ast.FunctionDeclaration) {
// 	fnName := function.GetEnvBindingName()
//
// 	if _, found := self.functions[fnName]; found {
// 		log.Errorf("Function already defined: %s - function will be ignored", function.Name)
// 		return
// 	}
//
// 	self.functions[fnName] = function
//
// 	// log.Debugf("SetFunction: %s(%s)", function.Name, fnName)
// }

// func (self *Environment) SetObject(object *ast.ObjectDeclaration) {
// 	if _, found := self.objects[object.Name.Name]; found {
// 		log.Errorf("Object already defined: %s - object will be ignored", object.Name.Name)
// 		return
// 	}
//
// 	self.objects[object.Name.Name] = object
//
// 	// log.Debugf("SetObject: %s", object.Name)
// }

// func (self *Environment) DefineCustomFunctionWithReceiver(receiver *ast.TypedIdentifier, name string, cb FunctionTypeCallback) *ast.FunctionDeclaration {
// 	fn := self.DefineCustomFunction(name, cb)
// 	fn.Receiver = receiver
// 	return fn
// }

// func (self *Environment) DefineCustomFunction(name string, cb FunctionTypeCallback) *ast.FunctionDeclaration {
// 	fn := &ast.FunctionDeclaration{
// 		Name:         name,
// 		Receiver:     nil,
// 		CustomFuncCb: cb,
// 	}
//
// 	self.SetFunction(fn)
//
// 	return fn
// }

// func (self *Environment) LookupObjectFunction(objectName, functionName string) *ast.FunctionDeclaration {
// 	object := self.LookupObject(objectName)
// 	if object == nil {
// 		panic("Object not found: " + objectName)
// 	}
//
// 	fn, ok := object.Methods[functionName]
// 	if !ok {
// 		panic("Method not found(" + objectName + "): " + functionName)
// 	}
//
// 	return fn
//
// }

func (self *Environment) GetRoot() *Environment {
	if self.parent == nil {
		return self
	}
	return self.parent.GetRoot()
}

// func (self *Environment) IsEvaluating(outer ast.Node, evalNode ast.Node) {
// 	self.evaluationState[evalNode.GetId()] = outer
// }

// func (self *Environment) SetEnum(t *ast.EnumDeclaration, evaluator *Evaluator) {
// 	if _, found := self.enums[t.Name.Name]; found {
// 		log.Errorf("Enum already defined: %s - enum will be ignored", t.Name.Name)
// 		return
// 	}
//
// 	self.enums[t.Name.Name] = t
//
// 	rtEnum := ast.NewRuntimeEnumDecl(t)
// 	self.SetVar(t.Name.Name, rtEnum)
// 	evaluator.evalEnumDeclaration(t, rtEnum)
// }
// func (self *Environment) LookupEnum(name string) *ast.EnumDeclaration {
// 	if t, ok := self.enums[name]; ok {
// 		return t
// 	}
// 	if self.parent != nil {
// 		return self.parent.LookupEnum(name)
// 	}
// 	return nil
// }
// func (self *Environment) LookupType(t string) ast.Type {
// 	if obj := self.LookupObject(t); obj != nil {
// 		return obj
// 	}
// 	if enum := self.LookupEnum(t); enum != nil {
// 		return enum
// 	}
// 	if bt, ok := ast.BasicTypes[t]; ok {
// 		return bt
// 	}
//
// 	if self.parent != nil {
// 		return self.parent.LookupType(t)
// 	}
//
// 	return nil
// }
