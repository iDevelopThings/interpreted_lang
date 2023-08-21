package ast

import (
	"errors"
	"sync/atomic"

	"github.com/charmbracelet/log"
)

var runtimeObjectUid = atomic.Int64{}

type RuntimeValueKind string

const (
	RuntimeValueKindUnknown RuntimeValueKind = "unknown"
	RuntimeValueKindObject  RuntimeValueKind = "object"
	RuntimeValueKindString  RuntimeValueKind = "string"
	RuntimeValueKindInteger RuntimeValueKind = "int"
	RuntimeValueKindFloat   RuntimeValueKind = "float"
	RuntimeValueKindBoolean RuntimeValueKind = "bool"
	RuntimeValueKindNull    RuntimeValueKind = "null"
	RuntimeValueKindArray   RuntimeValueKind = "array"
	RuntimeValueKindDict    RuntimeValueKind = "dict"

	// A global enum instance is bound to the runtime when the declaration is discovered
	// This allows for calling the `value constructors` like methods on the enum
	RuntimeValueKindEnumDecl  RuntimeValueKind = "enum"
	RuntimeValueKindEnumValue RuntimeValueKind = "enum_value"
)

type RuntimeValue struct {
	// Every runtime value has a unique id
	// This will allow us to track the value across the runtime
	// and compare values(in a simple manner)
	Uid int64

	TypeName string
	Decl     Declaration

	// This should be the raw value
	// For example, for an object instance, this would be the go map
	// For a string, this would be the go string
	// For a number, this would be the go int or float
	Value any

	Kind RuntimeValueKind

	Methods map[string]*FunctionDeclaration
}

func newRuntimeValue(decl Declaration) *RuntimeValue {
	rv := &RuntimeValue{
		Uid:     runtimeObjectUid.Add(1),
		Methods: map[string]*FunctionDeclaration{},
	}

	if decl != nil {
		rv.TypeName = decl.TypeName()
		rv.Decl = decl
	}

	return rv
}

func NewRuntimeArray(elementType Declaration) *RuntimeValue {
	rv := newRuntimeValue(elementType)

	rv.Value = []*RuntimeValue{}
	rv.Kind = RuntimeValueKindArray

	rv.Methods["length"] = &FunctionDeclaration{
		Name:       "length",
		ReturnType: NewIdentifierWithValue(nil, "int"),
		CustomFuncCb: func(args ...any) any {
			l := NewLiteral(nil, len(rv.Value.([]*RuntimeValue)))
			return NewRuntimeValueFromLiteral(
				l,
			)
		},
	}

	rv.Methods["add"] = &FunctionDeclaration{
		Name:       "add",
		ReturnType: NewIdentifierWithValue(nil, "void"),
		CustomFuncCb: func(args ...any) any {
			if len(args) == 0 {
				log.Fatalf("Expected >= 1 argument, got %d", len(args))
			}

			for _, arg := range args {
				rv.Value = append(rv.Value.([]*RuntimeValue), arg.(*RuntimeValue))
			}

			return nil
		},
	}

	rv.Methods["remove"] = &FunctionDeclaration{
		Name:       "remove",
		ReturnType: NewIdentifierWithValue(nil, "bool"),
		CustomFuncCb: func(args ...any) any {
			if len(args) == 0 {
				log.Fatalf("Expected >= 1 argument, got %d", len(args))
			}

			el := args[0].(*RuntimeValue)

			// Handle remove by index
			if el.Kind != rv.Kind && el.Kind == RuntimeValueKindInteger {
				index := el.Value.(int)
				if index < 0 || index >= len(rv.Value.([]*RuntimeValue)) {
					return false
				}

				rv.Value = append(rv.Value.([]*RuntimeValue)[:index], rv.Value.([]*RuntimeValue)[index+1:]...)
				return true
			}

			removed, err := rv.RemoveArrayElement(args[0].(*RuntimeValue))
			if err != nil {
				log.Fatalf("Cannot remove array element: %v", err)
			}

			return NewRuntimeValueFromLiteral(
				NewLiteral(nil, removed),
			)
		},
	}

	return rv
}

func NewRuntimeObject(decl Declaration) *RuntimeValue {
	rv := newRuntimeValue(decl)

	rv.Value = map[string]*RuntimeValue{}
	rv.Kind = RuntimeValueKindObject
	if decl != nil {
		rv.Methods = decl.GetMethods()
	} else {
		rv.Methods = map[string]*FunctionDeclaration{}
	}

	return rv
}

func NewRuntimeDictionary() *RuntimeValue {
	rv := newRuntimeValue(nil)

	rv.Value = map[string]*RuntimeValue{}
	rv.Kind = RuntimeValueKindDict
	rv.Methods = map[string]*FunctionDeclaration{}

	return rv
}

func NewRuntimeLiteral(value any) *RuntimeValue {
	rv := newRuntimeValue(nil)

	var lit *Literal
	switch value.(type) {
	case *Literal:
		lit = value.(*Literal)
	default:
		lit = NewLiteral(nil, value)
	}

	rv.TypeName = string(lit.Kind)
	rv.Value = lit.Value
	rv.Kind = RuntimeValueKind(lit.Kind)

	return rv
}

func NewRuntimeEnumDecl(enum *EnumDeclaration) *RuntimeValue {
	rv := newRuntimeValue(enum)

	rvValue := map[string]*RuntimeValue{}
	rv.Value = rvValue
	rv.Kind = RuntimeValueKindEnumDecl

	return rv
}

func NewRuntimeEnumValue(enum *EnumDeclaration) *RuntimeValue {
	rv := NewRuntimeDictionary()
	rv.Kind = RuntimeValueKindEnumValue

	return rv
}

func (self *RuntimeValue) GetField(fieldName string) *RuntimeValue {
	if self == nil {
		return nil
	}
	if self.Kind != RuntimeValueKindObject &&
		self.Kind != RuntimeValueKindDict &&
		self.Kind != RuntimeValueKindEnumValue &&
		self.Kind != RuntimeValueKindEnumDecl {
		return nil
	}

	fields, ok := self.Value.(map[string]*RuntimeValue)
	if !ok {
		return nil
	}

	if field, ok := fields[fieldName]; ok {
		return field
	}

	return nil
}
func (self *RuntimeValue) SetField(fieldName string, value *RuntimeValue) {
	if self.Kind != RuntimeValueKindObject &&
		self.Kind != RuntimeValueKindDict &&
		self.Kind != RuntimeValueKindEnumValue {
		return
	}

	if fields, ok := self.Value.(map[string]*RuntimeValue); ok {
		fields[fieldName] = value
	}
}
func (self *RuntimeValue) DeleteDictElement(name string) bool {
	if self.Kind != RuntimeValueKindDict {
		return false
	}

	if fields, ok := self.Value.(map[string]*RuntimeValue); ok {
		if field, ok := fields[name]; ok {
			field.Value = nil

			delete(fields, name)

			log.Debugf("Deleted dict element %s from %v", name, self.Value)

			return true
		}
	}

	return false
}

func (self *RuntimeValue) GetArrayElement(index int) *RuntimeValue {
	if self.Kind != RuntimeValueKindArray {
		return nil
	}

	if fields, ok := self.Value.([]*RuntimeValue); ok {
		return fields[index]
	}

	return nil
}
func (self *RuntimeValue) SetArrayElement(index int, value *RuntimeValue) {
	if self.Kind != RuntimeValueKindArray {
		return
	}

	if fields, ok := self.Value.([]*RuntimeValue); ok {
		fields[index] = value
	}
}
func (self *RuntimeValue) RemoveArrayElement(value ...*RuntimeValue) (bool, error) {
	if self.Kind != RuntimeValueKindArray {
		return false, nil
	}

	elements, ok := self.Value.([]*RuntimeValue)
	if !ok {
		return false, nil
	}

	for i, v := range elements {
		for _, runtimeValue := range value {
			if v == runtimeValue {
				elements = append(elements[:i], elements[i+1:]...)
				self.Value = elements
				return true, nil
			}
		}

	}

	return false, errors.New("element not found")
}
func (self *RuntimeValue) RemoveArrayElementsInRange(start, end int) (bool, error) {
	if self.Kind != RuntimeValueKindArray {
		return false, nil
	}

	elements, ok := self.Value.([]*RuntimeValue)
	if !ok {
		return false, nil
	}

	if start < 0 || start >= len(elements) {
		return false, errors.New("start index out of range")
	}

	if end < 0 || end >= len(elements) {
		return false, errors.New("end index out of range")
	}

	elements = append(elements[:start], elements[end+1:]...)

	self.Value = elements

	return true, nil
}
func (self *RuntimeValue) RemoveArrayElementByIndex(value *RuntimeValue) (bool, error) {
	if self.Kind != RuntimeValueKindArray {
		return false, nil
	}

	idx, ok := value.Value.(int)
	if !ok {
		return false, errors.New("array index must be an integer")
	}

	if elements, ok := self.Value.([]*RuntimeValue); ok {
		if idx >= 0 && idx < len(elements) {
			elements = append(elements[:idx], elements[idx+1:]...)
			self.Value = elements
			return true, nil
		}
	}

	return false, errors.New("element not found")
}

func (self *RuntimeValue) GetMethod(name string) *FunctionDeclaration {
	if method, ok := self.Methods[name]; ok {
		return method
	}
	return nil
}

func NewRuntimeValueFromLiteral(lit *Literal) *RuntimeValue {
	return &RuntimeValue{
		TypeName: string(lit.Kind),
		Value:    lit.Value,
		Kind:     RuntimeValueKind(lit.Kind),
	}
}

type ObjectFieldGetter interface {
	GetField(fieldName string) *RuntimeValue
	SetField(fieldName string, value *RuntimeValue)
}
