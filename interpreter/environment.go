package interpreter

import (
	"github.com/charmbracelet/log"

	"arc/ast"
)

type FunctionTypeCallback = func(args ...any) any

type HttpEnv struct {
	Options *ast.HttpServerConfig
	Routes  []*ast.HttpRouteDeclaration
}

type Environment struct {
	parent *Environment

	HttpEnv *HttpEnv

	vars    map[string]any
	objects map[string]*ast.ObjectDeclaration
	// objectInstances map[string]*ObjectInstantiation
	functions map[string]*ast.FunctionDeclaration
}

func NewEnvironment() *Environment {
	return &Environment{
		vars:    map[string]any{},
		objects: map[string]*ast.ObjectDeclaration{},
		// objectInstances: map[string]*ObjectInstantiation{},
		functions: map[string]*ast.FunctionDeclaration{},
		HttpEnv: &HttpEnv{
			Options: &ast.HttpServerConfig{
				Port:          ast.NewLiteral(nil, 8080),
				FormMaxMemory: ast.NewLiteral(nil, 10<<20), // 10 MB
			},
			Routes: make([]*ast.HttpRouteDeclaration, 0),
		},
	}
}

func (self *Environment) NewChild() *Environment {
	env := NewEnvironment()
	env.parent = self

	return env
}

func (self *Environment) LookupVar(name string) any {
	if v, ok := self.vars[name]; ok {
		return v
	}
	if self.parent != nil {
		return self.parent.LookupVar(name)
	}
	return nil
}

func (self *Environment) LookupFunction(name string) *ast.FunctionDeclaration {
	if f, ok := self.functions[name]; ok {
		return f
	}
	if self.parent != nil {
		return self.parent.LookupFunction(name)
	}
	return nil
}

func (self *Environment) LookupObject(name string) *ast.ObjectDeclaration {
	if o, ok := self.objects[name]; ok {
		return o
	}
	if self.parent != nil {
		return self.parent.LookupObject(name)
	}

	if bt, ok := ast.BasicTypes[name]; ok {
		return &ast.ObjectDeclaration{
			AstNode: bt.GetAstNode(),
			Name:    ast.NewIdentifierWithValue(nil, bt.TypeName()),
		}
	}

	return nil
}

func (self *Environment) SetVar(name string, value any) {
	self.vars[name] = value
}
func (self *Environment) AppendVars(vars map[string]any) {
	for k, v := range vars {
		self.vars[k] = v
	}
}
func (self *Environment) DeleteVar(name string) bool {
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

func (self *Environment) SetFunction(function *ast.FunctionDeclaration) {
	fnName := function.GetEnvName()

	if _, found := self.functions[fnName]; found {
		log.Errorf("Function already defined: %s - function will be ignored", function.Name)
		return
	}

	self.functions[fnName] = function

	// log.Debugf("SetFunction: %s(%s)", function.Name, fnName)
}

func (self *Environment) SetObject(object *ast.ObjectDeclaration) {
	if _, found := self.objects[object.Name.Name]; found {
		log.Errorf("Object already defined: %s - object will be ignored", object.Name.Name)
		return
	}

	self.objects[object.Name.Name] = object

	// log.Debugf("SetObject: %s", object.Name)
}

func (self *Environment) DefineCustomFunctionWithReceiver(receiver *ast.TypedIdentifier, name string, cb FunctionTypeCallback) *ast.FunctionDeclaration {
	fn := self.DefineCustomFunction(name, cb)
	fn.Receiver = receiver
	return fn
}
func (self *Environment) DefineCustomFunction(name string, cb FunctionTypeCallback) *ast.FunctionDeclaration {
	fn := &ast.FunctionDeclaration{
		Name:         name,
		Receiver:     nil,
		CustomFuncCb: cb,
	}

	self.SetFunction(fn)

	return fn
}

func (self *Environment) LookupObjectFunction(objectName, functionName string) *ast.FunctionDeclaration {
	object := self.LookupObject(objectName)
	if object == nil {
		panic("Object not found: " + objectName)
	}

	fn, ok := object.Methods[functionName]
	if !ok {
		panic("Method not found(" + objectName + "): " + functionName)
	}

	return fn

}

func (self *Environment) GetRoot() *Environment {
	if self.parent == nil {
		return self
	}
	return self.parent.GetRoot()
}

func (self *Environment) RegisterRoute(route *ast.HttpRouteDeclaration) {
	root := self.GetRoot()
	root.HttpEnv.Routes = append(root.HttpEnv.Routes, route)
}

func (self *Environment) SetHttpConfig(config *ast.HttpServerConfig) {
	root := self.GetRoot()
	root.HttpEnv.Options = config
}
