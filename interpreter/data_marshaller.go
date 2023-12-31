package interpreter

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/goccy/go-json"

	"arc/ast"
	"arc/log"
)

var (
	CannotMarshalRuntimeValueError = errors.New("cannot marshal runtime value")
)

func MarshalRuntimeObject(value *ast.RuntimeValue) map[string]any {
	jsonData := map[string]any{}
	if value.Kind != ast.RuntimeValueKindObject {
		log.Fatalf("Cannot marshal runtime object: %v", value)
	}
	if value.Decl == nil {
		log.Fatalf("Cannot marshal runtime object - declaration is missing: %v", value)
	}

	decl := value.Decl.(*ast.ObjectDeclaration)
	fields := value.Value.(map[string]*ast.RuntimeValue)

	for _, field := range decl.Fields {
		value, exists := fields[field.Name]
		if !exists {
			continue
		}

		if nestedDecl := Registry.LookupObject(field.TypeReference.Type); nestedDecl != nil && value.Kind == ast.RuntimeValueKindObject {
			jsonData[field.Name] = MarshalRuntimeObject(value)
		} else {
			jsonData[field.Name] = value.Value
		}
	}

	return jsonData
}

func MarshalRuntimeDictionary(value *ast.RuntimeValue) map[string]any {
	jsonData := map[string]any{}
	if value.Kind != ast.RuntimeValueKindDict {
		log.Fatalf("Cannot marshal runtime object: %v", value)
	}

	fields := value.Value.(map[string]*ast.RuntimeValue)

	for key, value := range fields {
		if value.Kind == ast.RuntimeValueKindDict {
			jsonData[key] = MarshalRuntimeDictionary(value)
		} else if nestedDecl := Registry.LookupObject(value.TypeName); nestedDecl != nil && value.Kind == ast.RuntimeValueKindObject {
			jsonData[key] = MarshalRuntimeObject(value)
		} else {
			jsonData[key] = value.Value
		}
	}

	return jsonData
}

func MarshalRuntimeValue(value *ast.RuntimeValue) ([]byte, error) {
	if value == nil {
		return nil, fmt.Errorf("%w: %v", CannotMarshalRuntimeValueError, value)
	}
	switch value.Kind {

	case ast.RuntimeValueKindObject:
		return json.Marshal(MarshalRuntimeObject(value))
	case ast.RuntimeValueKindArray:
		values := value.Value.([]*ast.RuntimeValue)
		jsonData := make([]any, len(values))
		for i, v := range values {
			jsonData[i] = v.Value
		}
		return json.Marshal(jsonData)

	case ast.RuntimeValueKindDict:
		return json.Marshal(MarshalRuntimeDictionary(value))
	default:
		return json.Marshal(value.Value)
	}

	// case *ast.ObjectRuntimeValue:
	// 	return json.Marshal(MarshalRuntimeObject(value, env))

	// case *ast.DictionaryRuntimeValue:
	// 	return json.Marshal(value)

	// return nil, fmt.Errorf("%w: %v", CannotMarshalRuntimeValueError, value)
}

func UnmarshalRuntimeValue(value any) (*ast.RuntimeValue, error) {
	if value == nil {
		return nil, fmt.Errorf("%w: %v", CannotMarshalRuntimeValueError, value)
	}

	if v, ok := value.(*ast.RuntimeValue); ok {
		return v, nil
	}

	switch value := value.(type) {
	case string, int, float64, bool, nil:
		return ast.NewRuntimeValueFromLiteral(
			ast.NewLiteral(nil, value),
		), nil

	case map[string]any:
		return UnmarshalRuntimeDictionary(value), nil
	case []any:
		return UnmarshalRuntimeArray(value), nil
	case url.Values:
		return UnmarshalUrlValues(value), nil

		// case map[string]any:
		// 	return UnmarshalRuntimeObject(env, value)
		//
		// case []any:
		// 	return UnmarshalRuntimeArray(env, value)
		//
	}

	return nil, fmt.Errorf("%w: %v", CannotMarshalRuntimeValueError, value)
}
func UnmarshalRuntimeValueAs(value []byte, typ ast.Type) (*ast.RuntimeValue, error) {
	if value == nil {
		return nil, fmt.Errorf("%w: %v", CannotMarshalRuntimeValueError, value)
	}

	// d := map[string]any{}
	var d any
	if err := json.Unmarshal(value, &d); err != nil {
		return nil, err
	}

	return UnmarshalRuntimeValue(d)

	// decl := Registry.LookupType(typ.TypeName())
	// if t, ok := ast.BasicTypes[typ.TypeName()]; ok {
	// 	decl = t
	// }

	/*switch typ := decl.(type) {
	case *ast.BasicType:
		switch typ.Name {
		case "dict":
			d := map[string]any{}
			if err := json.Unmarshal(value, &d); err != nil {
				return nil, err
			}
			return UnmarshalRuntimeDictionary(d), nil
		default:
			var d any
			if err := json.Unmarshal(value, &d); err != nil {
				return nil, err
			}
			return ast.NewRuntimeValueFromLiteral(ast.NewLiteral(nil, d)), nil
		}
	case *ast.ObjectDeclaration:
		d := map[string]any{}
		if err := json.Unmarshal(value, &d); err != nil {
			return nil, err
		}
		return UnmarshalRuntimeObject(typ, d), nil
	default:
		return nil, fmt.Errorf("cannot unmarshal runtime value as %s", typ.TypeName())
	}*/
}

func UnmarshalRuntimeArray(value []any) *ast.RuntimeValue {
	array := ast.NewRuntimeArray(nil)

	valueType := ast.RuntimeValueKindUnknown

	for _, value := range value {
		val, err := UnmarshalRuntimeValue(value)
		if err != nil {
			log.Warnf("cannot unmarshal array value - error: %s", err)
			continue
		}
		if valueType == ast.RuntimeValueKindUnknown {
			valueType = val.Kind
		}
		array.Value = append(array.Value.([]*ast.RuntimeValue), val)
	}

	decl := Registry.LookupObject(string(valueType))
	if decl != nil {
		array.TypeName = decl.TypeName()
		array.Decl = decl
	}

	return array
}

func UnmarshalRuntimeArrayWithDeclaration(decl *ast.ObjectDeclaration, value []*ast.RuntimeValue) *ast.RuntimeValue {
	array := ast.NewRuntimeArray(decl)

	for _, value := range value {
		val, err := UnmarshalRuntimeValue(value)
		if err != nil {
			log.Warnf("cannot unmarshal array value - error: %s", err)
			continue
		}

		array.Value = append(array.Value.([]*ast.RuntimeValue), val)
	}
	return array
}

func UnmarshalUrlValues(value url.Values) *ast.RuntimeValue {
	dict := ast.NewRuntimeDictionary()

	formValues := parseFormValues(value)

	fields := map[string]*ast.RuntimeValue{}

	for key, value := range formValues {
		val, err := UnmarshalRuntimeValue(value)
		if err != nil {
			log.Warnf("cannot unmarshal field %s - error: %s", key, err)
			continue
		}
		fields[key] = val
	}

	dict.Value = fields

	return dict
}

func UnmarshalRuntimeDictionary(value map[string]any) *ast.RuntimeValue {
	dict := ast.NewRuntimeDictionary()

	fields := map[string]*ast.RuntimeValue{}

	for key, value := range value {
		val, err := UnmarshalRuntimeValue(value)
		if err != nil {
			log.Warnf("cannot unmarshal field %s - error: %s", key, err)
			continue
		}
		fields[key] = val
	}

	dict.Value = fields

	return dict
}

func UnmarshalRuntimeObject(decl *ast.ObjectDeclaration, value map[string]any) *ast.RuntimeValue {
	obj := ast.NewRuntimeObject(decl)

	fields := map[string]*ast.RuntimeValue{}

	for _, field := range decl.Fields {
		value, exists := value[field.Name]
		if !exists {
			continue
		}

		if field.TypeReference.Type == "" {
			log.Warnf("cannot unmarshal field %s of type %s", field.Name, field.TypeReference.Type)
			continue
		}

		// If we have a basic type, it's likely a literal value assign
		if bt := field.TypeReference.GetBasicType(); bt != nil {
			val, err := UnmarshalRuntimeValue(value)
			if err != nil {
				log.Warnf("cannot unmarshal field %s of type %s - error: %s", field.Name, field.TypeReference.Type, err)
				continue
			}
			fields[field.Name] = val
		} else if nestedDecl := Registry.LookupObject(field.TypeReference.Type); nestedDecl != nil {
			m, ok := value.(map[string]any)
			if !ok {
				log.Warnf("cannot unmarshal field %s of type %s - expected object", field.Name, field.TypeReference.Type)
				continue
			}
			fields[field.Name] = UnmarshalRuntimeObject(nestedDecl, m)
		} else {
			val, err := UnmarshalRuntimeValue(value)
			if err != nil {
				log.Warnf("cannot unmarshal field %s of type %s - error: %s", field.Name, field.TypeReference.Type, err)
				continue
			}
			fields[field.Name] = val
		}
	}

	obj.Value = fields

	return obj
}

func UnmarshalRuntimeObjectFromDictionary(decl *ast.ObjectDeclaration, rtValue *ast.RuntimeValue) *ast.RuntimeValue {
	obj := ast.NewRuntimeObject(decl)

	fields := map[string]*ast.RuntimeValue{}

	for _, field := range decl.Fields {
		value := rtValue.GetField(field.Name)
		if value == nil {
			continue
		}

		if field.TypeReference.Type == "" {
			log.Warnf("cannot unmarshal field %s of type %s", field.Name, field.TypeReference.Type)
			continue
		}

		/* else if value.Kind == "dict" {
			nestedDecl := env.LookupObject(field.Type)
			if nestedDecl != nil {
				fields[field.Name] = UnmarshalRuntimeObjectFromDictionary(nestedDecl, env, value)
			} else {
				log.Warnf("cannot unmarshal dict field %s of type %s", field.Name, field.TypeReference.Type)
				continue
			}
		}*/

		if nestedDecl := Registry.LookupObject(field.TypeReference.Type); nestedDecl != nil {
			if _, ok := value.Value.(map[string]*ast.RuntimeValue); ok {
				fields[field.Name] = UnmarshalRuntimeObjectFromDictionary(nestedDecl, value)
			} else if d, ok := value.Value.([]*ast.RuntimeValue); ok {
				fields[field.Name] = UnmarshalRuntimeArrayWithDeclaration(nestedDecl, d)
			} else if d, ok := value.Value.(map[string]any); ok {
				fields[field.Name] = UnmarshalRuntimeObject(nestedDecl, d)
			} else {
				val, err := UnmarshalRuntimeValue(value)
				if err != nil {
					log.Warnf("cannot unmarshal field %s of type %s - error: %s", field.Name, field.TypeReference.Type, err)
					continue
				}
				fields[field.Name] = val
			}
		} else {
			val, err := UnmarshalRuntimeValue(value)
			if err != nil {
				log.Warnf("cannot unmarshal field %s of type %s - error: %s", field.Name, field.TypeReference.Type, err)
				continue
			}
			fields[field.Name] = val
		}
	}

	obj.Value = fields

	return obj
}

func parseFormValues(values url.Values) map[string]any {
	nested := make(map[string]any)

	for key, values := range values {
		parts := strings.Split(key, "[")
		if len(parts) < 2 || parts[1] == "]" {
			key = strings.TrimSuffix(key, "[]")
			nested[key] = []any{}
			for _, value := range values {
				nested[key] = append(nested[key].([]any), value)
			}
			continue
		}

		inner := nested
		for i, part := range parts {
			part = strings.TrimRight(part, "]")

			if i == len(parts)-1 {
				inner[part] = values[0]
			} else {
				if _, exists := inner[part]; !exists {
					inner[part] = make(map[string]any)
				}
				inner = inner[part].(map[string]any)
			}
		}
	}
	return nested
}
