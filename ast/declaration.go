package ast

import (
	"arc/log"
)

type Declaration interface {
	Node
	IsDeclaration()
	TypeName() string
	GetMethods() map[string]*FunctionDeclaration
}

type ObjectDeclaration struct {
	*AstNode
	Name     *Identifier
	Fields   []*TypedIdentifier
	Methods  map[string]*FunctionDeclaration
	IsExtern bool
}

func (self *ObjectDeclaration) IsTopLevelStatement()      {}
func (self *ObjectDeclaration) IsStatement()              {}
func (self *ObjectDeclaration) IsDeclaration()            {}
func (self *ObjectDeclaration) TypeName() string          { return self.Name.Name }
func (self *ObjectDeclaration) GetEnvBindingName() string { return self.Name.Name }

func (self *ObjectDeclaration) GetMethods() map[string]*FunctionDeclaration {
	return self.Methods
}

func (self *ObjectDeclaration) GetMethod(name string) *FunctionDeclaration {
	if m, ok := self.Methods[name]; ok {
		return m
	}
	return nil
}

func (self *ObjectDeclaration) Merge(object *ObjectDeclaration) {
	for _, field := range object.Fields {
		hasExisting := false
		for _, identifier := range self.Fields {
			if identifier.Name == field.Name {
				log.Warnf("Field %s already exists in object %s", identifier.Name, self.Name.Name)
				hasExisting = true
				continue
			}
		}
		if !hasExisting {
			self.Fields = append(self.Fields, field)
		}
	}

	for _, method := range object.Methods {
		if self.Methods == nil {
			self.Methods = make(map[string]*FunctionDeclaration)
		}
		if _, ok := self.Methods[method.Name]; ok {
			log.Warnf("Method %s already exists in object %s", method.Name, self.Name.Name)
			continue
		}
		self.Methods[method.Name] = method
	}
}

type FunctionDeclaration struct {
	*AstNode
	Name            string
	Args            []*TypedIdentifier
	ReturnType      *TypeReference
	Receiver        *TypedIdentifier
	Body            *Block
	CustomFuncCb    func(args ...any) any `json:"-"`
	IsStatic        bool
	IsBuiltin       bool
	IsAnonymous     bool
	IsExtern        bool
	HasVariadicArgs bool
}

func (self *FunctionDeclaration) IsTopLevelStatement()                        {}
func (self *FunctionDeclaration) IsStatement()                                {}
func (self *FunctionDeclaration) IsDeclaration()                              {}
func (self *FunctionDeclaration) GetMethods() map[string]*FunctionDeclaration { return nil }
func (self *FunctionDeclaration) TypeName() string                            { return self.Name }
func (self *FunctionDeclaration) GetEnvBindingName() string {
	if self.Receiver != nil {
		return self.Receiver.TypeReference.Type + "_" + self.Name
	}
	return self.Name
}

type EnumDeclaration struct {
	*AstNode
	Name   *Identifier
	Values []*EnumValue
}

func (self *EnumDeclaration) IsTopLevelStatement()                        {}
func (self *EnumDeclaration) IsStatement()                                {}
func (self *EnumDeclaration) IsDeclaration()                              {}
func (self *EnumDeclaration) GetMethods() map[string]*FunctionDeclaration { return nil }
func (self *EnumDeclaration) TypeName() string                            { return self.Name.Name }
func (self *EnumDeclaration) GetEnvBindingName() string                   { return self.Name.Name }

func (self *EnumDeclaration) GetValueConstructor(name string) *EnumValue {
	for _, value := range self.Values {
		if value.Name.Name == name && value.Kind == EnumValueKindWithValue {
			return value
		}
	}
	return nil
}
func (self *EnumDeclaration) HasValueConstructor(name string) bool {
	return self.GetValueConstructor(name) != nil
}

type EnumValueKind string

const (
	EnumValueKindLiteral   EnumValueKind = "literal"
	EnumValueKindWithValue EnumValueKind = "with_value"
)

type EnumValue struct {
	*AstNode
	Name *Identifier
	Kind EnumValueKind
	// We only have a value & type when we're using an `EnumValueKindLiteral`
	Type  Type
	Value Expr
	// We only have properties when we're using an `EnumValueKindWithValue`
	Properties []*TypedIdentifier
}
