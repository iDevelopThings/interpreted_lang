package interpreter

import (
	"github.com/charmbracelet/log"

	"arc/ast"
)

var Registry *RegistryInstance = NewRegistry()

func NewRegistry() *RegistryInstance {
	inst := &RegistryInstance{
		objects:   map[string]RegistryItem[*ast.ObjectDeclaration]{},
		enums:     map[string]RegistryItem[*ast.EnumDeclaration]{},
		functions: map[string]RegistryItem[*ast.FunctionDeclaration]{},
		HttpEnv: &HttpEnv{
			Routes: make([]*ast.HttpRouteDeclaration, 0),
		},
	}

	return inst
}

func init() {
	for name, t := range ast.BasicTypes {
		Registry.SetObject(&ast.ObjectDeclaration{
			AstNode: t.GetAstNode(),
			Name:    ast.NewIdentifierWithValue(nil, name),
		})
	}
}

type HttpEnv struct {
	Routes []*ast.HttpRouteDeclaration
}

type RegistryItem[T any] struct {
	SourceFile  string
	Declaration T
}

type RegistryInstance struct {
	HttpEnv *HttpEnv

	objects   map[string]RegistryItem[*ast.ObjectDeclaration]
	functions map[string]RegistryItem[*ast.FunctionDeclaration]
	enums     map[string]RegistryItem[*ast.EnumDeclaration]
}

func (self *RegistryInstance) hi() {

}

func (self *RegistryInstance) LookupFunction(name string) *ast.FunctionDeclaration {
	if f, ok := self.functions[name]; ok {
		return f.Declaration
	}
	return nil
}

func (self *RegistryInstance) LookupObject(name string) *ast.ObjectDeclaration {
	if o, ok := self.objects[name]; ok {
		return o.Declaration
	}
	return nil
}

func (self *RegistryInstance) SetFunction(function *ast.FunctionDeclaration) {
	fnName := function.GetEnvBindingName()

	if _, found := self.functions[fnName]; found {
		log.Errorf("Function already defined: %s - function will be ignored", function.Name)
		return
	}

	sf := ""
	if function.AstNode != nil && function.AstNode.Token != nil {
		sf = function.AstNode.Token.Source
	}

	self.functions[fnName] = RegistryItem[*ast.FunctionDeclaration]{
		SourceFile:  sf,
		Declaration: function,
	}
}

func (self *RegistryInstance) SetObject(object *ast.ObjectDeclaration) {
	if _, found := self.objects[object.Name.Name]; found {
		log.Errorf("Object already defined: %s - object will be ignored", object.Name.Name)
		return
	}

	sf := object.AstNode.Token.Source
	// sf := ""
	// if object.AstNode != nil && object.AstNode.Token != nil {
	// 	sf = object.AstNode.Token.Source
	// } else {
	// 	print("")
	// }

	self.objects[object.Name.Name] = RegistryItem[*ast.ObjectDeclaration]{
		SourceFile:  sf,
		Declaration: object,
	}
}

func (self *RegistryInstance) SetEnum(t *ast.EnumDeclaration) *ast.RuntimeValue {
	if _, found := self.enums[t.Name.Name]; found {
		log.Errorf("Enum already defined: %s - enum will be ignored", t.Name.Name)
		return nil
	}

	self.enums[t.Name.Name] = RegistryItem[*ast.EnumDeclaration]{
		SourceFile:  t.AstNode.Token.Source,
		Declaration: t,
	}

	rtEnum := ast.NewRuntimeEnumDecl(t)

	return rtEnum
}

func (self *RegistryInstance) DefineCustomFunctionWithReceiver(receiver *ast.TypedIdentifier, name string, cb FunctionTypeCallback) *ast.FunctionDeclaration {
	fn := self.DefineCustomFunction(name, cb)
	fn.Receiver = receiver
	return fn
}

func (self *RegistryInstance) DefineCustomFunction(name string, cb FunctionTypeCallback) *ast.FunctionDeclaration {
	fn := &ast.FunctionDeclaration{
		Name:         name,
		Receiver:     nil,
		CustomFuncCb: cb,
	}

	self.SetFunction(fn)

	return fn
}

func (self *RegistryInstance) LookupObjectFunction(objectName, functionName string) *ast.FunctionDeclaration {
	object := self.LookupObject(objectName)
	if object == nil {
		panic("Object not found: " + objectName)
	}

	if fn, ok := object.Methods[functionName]; ok {
		return fn
	}

	return nil
}

func (self *RegistryInstance) LookupEnum(name string) *ast.EnumDeclaration {
	if t, ok := self.enums[name]; ok {
		return t.Declaration
	}
	return nil
}

func (self *RegistryInstance) RegisterRoute(route *ast.HttpRouteDeclaration) {
	self.HttpEnv.Routes = append(self.HttpEnv.Routes, route)
	log.Debugf("Registered route: %s %s", route.Method, route.Path.Value)
}

func (self *RegistryInstance) LookupType(t string) ast.Type {
	if obj := self.LookupObject(t); obj != nil {
		return obj
	}
	if enum := self.LookupEnum(t); enum != nil {
		return enum
	}
	return nil
}
