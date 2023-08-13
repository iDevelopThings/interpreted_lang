package interpreter

import (
	"errors"
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/goccy/go-json"

	"interpreted_lang/ast"
)

var (
	CannotMarshalRuntimeValueError = errors.New("cannot marshal runtime value")
)

func MarshalRuntimeObject(value *ast.RuntimeValue, env *Environment) map[string]any {
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

		if nestedDecl := env.LookupObject(field.Type); nestedDecl != nil && value.Kind == ast.RuntimeValueKindObject {
			jsonData[field.Name] = MarshalRuntimeObject(value, env)
		} else {
			jsonData[field.Name] = value.Value
		}
	}

	return jsonData
}

func MarshalRuntimeDictionary(value *ast.RuntimeValue, env *Environment) map[string]any {
	jsonData := map[string]any{}
	if value.Kind != ast.RuntimeValueKindDict {
		log.Fatalf("Cannot marshal runtime object: %v", value)
	}

	fields := value.Value.(map[string]*ast.RuntimeValue)

	for key, value := range fields {
		if nestedDecl := env.LookupObject(value.TypeName); nestedDecl != nil && value.Kind == ast.RuntimeValueKindObject {
			jsonData[key] = MarshalRuntimeObject(value, env)
		} else {
			jsonData[key] = value.Value
		}
	}

	return jsonData
}

func MarshalRuntimeValue(env *Environment, value *ast.RuntimeValue) ([]byte, error) {
	if value == nil {
		return nil, fmt.Errorf("%w: %v", CannotMarshalRuntimeValueError, value)
	}
	switch value.Kind {

	case ast.RuntimeValueKindObject:
		return json.Marshal(MarshalRuntimeObject(value, env))
	case ast.RuntimeValueKindArray:
		values := value.Value.([]*ast.RuntimeValue)
		jsonData := make([]any, len(values))
		for i, v := range values {
			jsonData[i] = v.Value
		}
		return json.Marshal(jsonData)

	case ast.RuntimeValueKindDict:
		return json.Marshal(MarshalRuntimeDictionary(value, env))
	default:
		return json.Marshal(value.Value)
	}

	// case *ast.ObjectRuntimeValue:
	// 	return json.Marshal(MarshalRuntimeObject(value, env))

	// case *ast.DictionaryRuntimeValue:
	// 	return json.Marshal(value)

	// return nil, fmt.Errorf("%w: %v", CannotMarshalRuntimeValueError, value)
}

func UnmarshalRuntimeValue(env *Environment, value any) (*ast.RuntimeValue, error) {
	switch value := value.(type) {

	case string, int, float64, bool, nil:
		return ast.NewRuntimeValueFromLiteral(
			ast.NewLiteral(nil, value),
		), nil

	case map[string]any:
		return UnmarshalRuntimeDictionary(env, value), nil

		// case map[string]any:
		// 	return UnmarshalRuntimeObject(env, value)
		//
		// case []any:
		// 	return UnmarshalRuntimeArray(env, value)
		//
	}

	return nil, fmt.Errorf("%w: %v", CannotMarshalRuntimeValueError, value)
}

func UnmarshalRuntimeDictionary(env *Environment, value map[string]any) *ast.RuntimeValue {
	dict := ast.NewRuntimeDictionary()

	fields := map[string]*ast.RuntimeValue{}

	for key, value := range value {
		val, err := UnmarshalRuntimeValue(env, value)
		if err != nil {
			log.Warnf("cannot unmarshal field %s - error: %s", key, err)
			continue
		}
		fields[key] = val
	}

	dict.Value = fields

	return dict
}

func UnmarshalRuntimeObject(decl *ast.ObjectDeclaration, env *Environment, value map[string]any) *ast.RuntimeValue {
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
		if nestedDecl := env.LookupObject(field.Type); nestedDecl != nil {
			fields[field.Name] = UnmarshalRuntimeObject(nestedDecl, env, value.(map[string]any))
		} else {
			val, err := UnmarshalRuntimeValue(env, value)
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
