package ast

import (
	"errors"
	"maps"
	"net/http"
	"sync/atomic"

	"arc/http_server"
	"arc/log"
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
	RuntimeValueKindNone    RuntimeValueKind = "none"
	RuntimeValueKindArray   RuntimeValueKind = "array"
	RuntimeValueKindDict    RuntimeValueKind = "dict"

	// Option kind is mainly an underlying runtime value wrapped in another...
	RuntimeValueKindOption RuntimeValueKind = "option"

	// A global enum instance is bound to the runtime when the declaration is discovered
	// This allows for calling the `value constructors` like methods on the enum
	RuntimeValueKindEnumDecl  RuntimeValueKind = "enum"
	RuntimeValueKindEnumValue RuntimeValueKind = "enum_value"

	// The wrapper object that we throw all of the main request data into when handling http requests
	// This is the object bound to the environment & accessible in the language route handler
	RuntimeValueKindRequestObject RuntimeValueKind = "request_object"
)

type RuntimeValue struct {
	// Every runtime value has a unique id
	// This will allow us to track the value across the runtime
	// and compare values(in a simple manner)
	Uid          int64
	OriginalNode Node

	// When we unwrap an option type value, we'll add it's outer here, so we can keep a ref
	UnwrapParent *RuntimeValue

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

func newRuntimeValue(node Node) *RuntimeValue {
	rv := &RuntimeValue{
		Uid:          runtimeObjectUid.Add(1),
		OriginalNode: node,
		Methods:      map[string]*FunctionDeclaration{},
	}

	if node != nil {
		if decl, ok := node.(Declaration); ok {
			rv.TypeName = decl.TypeName()
			rv.Decl = decl
		}
	}

	return rv
}

func NewRuntimeValueClone(original *RuntimeValue) *RuntimeValue {
	rv := newRuntimeValue(original.OriginalNode)

	if rv.Decl != original.Decl {
		rv.Decl = original.Decl
	}
	rv.TypeName = original.TypeName
	rv.Value = original.Value
	rv.Kind = original.Kind
	rv.Methods = maps.Clone(original.Methods)

	return rv
}

func NewRuntimeArray(elementType Declaration) *RuntimeValue {
	rv := newRuntimeValue(elementType)

	rv.Value = []*RuntimeValue{}
	rv.Kind = RuntimeValueKindArray

	rv.Methods["length"] = &FunctionDeclaration{
		Name:       "length",
		ReturnType: NewTypeReferenceWithValue("int"),
		CustomFuncCb: func(args ...any) any {
			l := NewLiteral(nil, len(rv.Value.([]*RuntimeValue)))
			return NewRuntimeValueFromLiteral(
				l,
			)
		},
	}

	rv.Methods["add"] = &FunctionDeclaration{
		Name:       "add",
		ReturnType: NewTypeReferenceWithValue("void"),
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
		ReturnType: NewTypeReferenceWithValue("bool"),
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

	rv.TypeName = "dict"
	rv.Value = map[string]*RuntimeValue{}
	rv.Kind = RuntimeValueKindDict
	rv.Methods = map[string]*FunctionDeclaration{}

	return rv
}

func NewRuntimeValueFromLiteral(lit *Literal) *RuntimeValue {
	rv := newRuntimeValue(lit)
	rv.TypeName = string(lit.Kind)
	rv.Value = lit.Value
	rv.Kind = RuntimeValueKind(lit.Kind)
	return rv
}

func NewRuntimeLiteral(value any) *RuntimeValue {
	rv := newRuntimeValue(nil)

	var lit *Literal
	switch value.(type) {
	case *Literal:
		lit = value.(*Literal)
		rv.OriginalNode = lit
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

func NewRuntimeOptionValue(value *RuntimeValue) *RuntimeValue {
	rv := newRuntimeValue(value.OriginalNode)

	rv.Value = value
	rv.Kind = RuntimeValueKindOption
	rv.TypeName = value.TypeName
	rv.Decl = value.Decl

	return rv
}

func NewRuntimeRequestObject(route *HttpRouteDeclaration, req *http.Request, res http.ResponseWriter, params http_server.Params) *RuntimeValue {
	rv := newRuntimeValue(route)

	rv.Kind = RuntimeValueKindRequestObject
	rv.TypeName = "HttpRequestWrapper"

	requestObj := NewRuntimeObject(nil)
	requestObj.TypeName = "HttpRequest"

	resObj := NewRuntimeObject(nil)
	resObj.TypeName = "HttpResponse"
	resObj.Value = res

	rv.Value = map[string]any{
		"internal_request":  req,
		"internal_response": resObj,
		"params":            params,
		"request":           requestObj,
	}

	rv.Methods = map[string]*FunctionDeclaration{}

	return rv
}

func RuntimeValueAs[T any](rv *RuntimeValue) T {
	value, ok := rv.Value.(T)
	if !ok {
		log.Fatalf("Cannot cast runtime value to type %T", value)
	}
	return value
}

func (self *RuntimeValue) GetField(fieldName string) *RuntimeValue {
	if self == nil {
		return nil
	}
	if self.Kind != RuntimeValueKindObject &&
		self.Kind != RuntimeValueKindDict &&
		self.Kind != RuntimeValueKindEnumValue &&
		self.Kind != RuntimeValueKindEnumDecl &&
		self.Kind != RuntimeValueKindRequestObject {
		return nil
	}

	switch v := self.Value.(type) {
	case map[string]*RuntimeValue:
		if field, ok := v[fieldName]; ok {
			return field
		}
	case map[string]any:
		if field, ok := v[fieldName]; ok {
			if v, ok := field.(*RuntimeValue); ok {
				return v
			}
		}
	}

	return nil
}
func (self *RuntimeValue) GetFieldValue(fieldName string) any {
	if self == nil {
		return nil
	}
	if self.Kind != RuntimeValueKindObject &&
		self.Kind != RuntimeValueKindDict &&
		self.Kind != RuntimeValueKindEnumValue &&
		self.Kind != RuntimeValueKindEnumDecl &&
		self.Kind != RuntimeValueKindRequestObject {
		return nil
	}

	switch v := self.Value.(type) {
	case map[string]*RuntimeValue:
		if field, ok := v[fieldName]; ok {
			return field
		}
	case map[string]any:
		if field, ok := v[fieldName]; ok {
			return field
		}
	default:
		log.Warnf("Unhandled field value type from `GetFieldValue`: %T - field: %s, struct: %#v", v, fieldName, self)
	}

	return nil
}
func (self *RuntimeValue) SetField(fieldName string, value *RuntimeValue) {
	if self.Kind != RuntimeValueKindObject &&
		self.Kind != RuntimeValueKindDict &&
		self.Kind != RuntimeValueKindEnumValue &&
		self.Kind != RuntimeValueKindRequestObject {
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

func (self *RuntimeValue) IsNumeric() bool {
	return self.Kind == RuntimeValueKindInteger || self.Kind == RuntimeValueKindFloat
}

func (self *RuntimeValue) Apply(value *RuntimeValue) {
	if value.Kind != self.Kind {
		log.Fatalf("Cannot apply runtime value of type %s to type %s", value.Kind, self.Kind)
	}

	self.Value = value.Value
}

// Mainly used for option types
// Since our main underlying value is wrapped in a runtime value with an option kind...
// We need to unwrap it to get the underlying value... but only if exists
func (self *RuntimeValue) HasValue() bool {
	if !self.IsOptionKind() {
		return self.Value != nil
	}

	// An option kind can be in two states:
	// 1. It has a `.value` with a rt value of `ast.NoneType` wrapped in a rt value
	// 2. It has a `.value` of nil(nil represents none) in value form at-least

	if self.Value == nil {
		return false
	}

	v, ok := self.Value.(*RuntimeValue)
	if !ok {
		return false
	}

	return !v.IsNoneKind()
}

func (self *RuntimeValue) Unwrap() *RuntimeValue {
	if self.Kind != RuntimeValueKindOption {
		return self
	}

	if self.Value == nil {
		return nil
	}

	v := self.Value.(*RuntimeValue)
	if v.UnwrapParent == nil {
		v.UnwrapParent = self
	}

	return v
}

func (self *RuntimeValue) IsUnknownKind() bool { return self.Kind == RuntimeValueKindUnknown }
func (self *RuntimeValue) IsObjectKind() bool  { return self.Kind == RuntimeValueKindObject }
func (self *RuntimeValue) IsStringKind() bool  { return self.Kind == RuntimeValueKindString }
func (self *RuntimeValue) IsIntegerKind() bool { return self.Kind == RuntimeValueKindInteger }
func (self *RuntimeValue) IsFloatKind() bool   { return self.Kind == RuntimeValueKindFloat }
func (self *RuntimeValue) IsBooleanKind() bool { return self.Kind == RuntimeValueKindBoolean }
func (self *RuntimeValue) IsNoneKind() bool    { return self.Kind == RuntimeValueKindNone }
func (self *RuntimeValue) IsArrayKind() bool   { return self.Kind == RuntimeValueKindArray }
func (self *RuntimeValue) IsDictKind() bool    { return self.Kind == RuntimeValueKindDict }
func (self *RuntimeValue) IsOptionKind(checkParent ...bool) bool {
	if len(checkParent) == 0 {
		checkParent = []bool{false}
	}

	if self.Kind == RuntimeValueKindOption {
		return true
	}

	if checkParent[0] && self.UnwrapParent != nil {
		return self.UnwrapParent.IsOptionKind()
	}

	return false
}
func (self *RuntimeValue) IsEnumDeclKind() bool  { return self.Kind == RuntimeValueKindEnumDecl }
func (self *RuntimeValue) IsEnumValueKind() bool { return self.Kind == RuntimeValueKindEnumValue }

type ObjectFieldGetter interface {
	GetField(fieldName string) *RuntimeValue
	SetField(fieldName string, value *RuntimeValue)
}
