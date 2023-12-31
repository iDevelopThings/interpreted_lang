package interpreter

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	gohttpclient "github.com/bozd4g/go-http-client"

	"arc/ast"
	"arc/log"
	"arc/utilities"
)

func getValueArgIndexes(fmtString string) []int {
	indexes := make([]int, 0)
	argIdxCounter := 0
	for i, r := range fmtString {
		if r == '%' {
			indexes = append(indexes, i)
			argIdxCounter++
		}
	}
	return indexes
}

//goland:noinspection GoSnakeCaseUsage
type RT_string struct{}

//goland:noinspection GoSnakeCaseUsage
func (t *RT_string) Static_to(env *Environment, args ...any) (interface{}, string) {
	return TypeCoercion.MustCast(ast.NewRuntimeLiteral(args[0]), ast.StringType), ""
}

//goland:noinspection GoSnakeCaseUsage
type RT_fmt struct {
}

//goland:noinspection GoSnakeCaseUsage
func (t *RT_fmt) Static_printf(env *Environment, fmtString string, args ...any) interface{} {
	valIndexes := getValueArgIndexes(fmtString)
	var logArgs []any
	for i, arg := range args {
		if i >= len(valIndexes) {
			break
		}
		argIdx := valIndexes[i]
		fmtArgType := fmtString[argIdx : argIdx+2]

		switch fmtArgType {
		case "%v":
			val := formatLogRuntimeValue(arg)
			logArgs = append(logArgs, val)

		default:
			logArgs = append(logArgs, arg)
		}

	}
	fmt.Printf(fmtString, logArgs...)
	return nil
}

//goland:noinspection GoSnakeCaseUsage
func (t *RT_fmt) Static_println(env *Environment, args ...any) interface{} {
	fmt.Println(args...)
	return nil
}

//goland:noinspection GoSnakeCaseUsage
func (t *RT_fmt) Static_print(env *Environment, args ...any) interface{} {
	fmt.Print(args...)
	return nil
}

//goland:noinspection GoSnakeCaseUsage
type RT_error struct{}

//goland:noinspection GoSnakeCaseUsage
func (t *RT_error) Static_panic(env *Environment, args ...any) interface{} {
	panic(args[0])
}

func bindStaticRuntimeType(t any) {
	rPtrType := reflect.TypeOf(t)
	rType := rPtrType.Elem()

	rVal := reflect.ValueOf(t)

	objName := rType.Name()
	objName = strings.Replace(objName, "RT_", "", 1)

	// env.SetObject()
	objDecl := &ast.ObjectDeclaration{
		AstNode: ast.NewAstNode(nil),
		Name:    ast.NewIdentifierWithValue(nil, objName),
		Fields:  make([]*ast.TypedIdentifier, 0),
		Methods: make(map[string]*ast.FunctionDeclaration),
	}

	for i := 0; i < rType.NumField(); i++ {
		f := rType.Field(i)
		objDecl.Fields = append(objDecl.Fields, &ast.TypedIdentifier{
			Identifier: ast.NewIdentifierWithValue(nil, f.Name),
			TypeReference: &ast.TypeReference{
				AstNode: ast.NewAstNode(nil),
				Type:    f.Type.Name(),
			},
		})
	}

	for i := 0; i < rPtrType.NumMethod(); i++ {
		m := rPtrType.Method(i)
		mVal := rVal.Method(i)

		methodDecl := &ast.FunctionDeclaration{
			AstNode:      ast.NewAstNode(nil),
			Name:         m.Name,
			Args:         nil,
			CustomFuncCb: nil,
			IsBuiltin:    true,
		}

		if strings.HasPrefix(m.Name, "Static_") {
			methodDecl.IsStatic = true
			methodDecl.Name = strings.TrimPrefix(m.Name, "Static_")

			methodDecl.Receiver = &ast.TypedIdentifier{
				TypeReference: &ast.TypeReference{
					Type: objName,
				},
			}
		} else {
			panic("not implemented")
		}

		if m.Type.NumOut() > 1 {
			// Output types are in groups of two
			// any{} & go type name
			// only so we can convert it to a runtime value with type

			for i := 0; i < m.Type.NumOut(); i += 2 {
				ident := ast.NewIdentifierWithValue(nil, m.Type.Out(i+1).Name())
				methodDecl.ReturnType = &ast.TypeReference{
					AstNode: ast.NewAstNode(nil),
					Type:    ident.Name,
				}
			}

		}

		if argsC := m.Type.NumIn(); argsC > 1 {
			methodDecl.Args = make([]*ast.TypedIdentifier, 0)
			for argIdx := 1; argIdx < argsC; argIdx++ {
				arg := m.Type.In(argIdx)
				if argIdx == 1 {
					if arg.String() != "*interpreter.Environment" {
						log.Fatalf("First argument of a runtime bound function must be *interpreter.Environment, got %s", arg.String())
					}
					// We don't want to add the env arg to the declaration info list
					continue
				}

				name := arg.String()
				isVariadic := argIdx == argsC-1 && m.Type.IsVariadic()
				if isVariadic {
					// Extract the actual type of the variadic arguments
					name = arg.Elem().String()
				}

				if name == "interface {}" {
					name = "any"
				}

				methodDecl.Args = append(methodDecl.Args, &ast.TypedIdentifier{
					Identifier: ast.NewIdentifierWithValue(nil, fmt.Sprintf("arg%d", argIdx)),
					TypeReference: &ast.TypeReference{
						AstNode:    ast.NewAstNode(nil),
						Type:       name,
						IsVariadic: isVariadic,
					},
				})

				methodDecl.HasVariadicArgs = isVariadic
			}
		}

		if methodDecl.IsStatic {
			methodDecl.CustomFuncCb = func(args ...interface{}) interface{} {
				// timer := utilities.NewTimer("runtime bound function execution " + objName + "." + methodDecl.Name)
				// defer timer.StopAndLog()

				fnArgValues := make([]reflect.Value, 0)
				// fnArgValues = append(fnArgValues, reflect.ValueOf(t))

				isVariadicInsertion := false

				var mArg *ast.TypedIdentifier
				for idx := 1; idx < len(args); idx++ {
					arg := args[idx]
					if mArg == nil || !mArg.TypeReference.IsVariadic {
						mArg = methodDecl.Args[idx-1]
					}
					if mArg == nil {
						panic("mArg is nil")
					}
					if mArg.TypeReference.IsVariadic {
						isVariadicInsertion = true
					}

					if bt := mArg.TypeReference.GetBasicType(); bt != nil {
						value := arg
						if value == (*ast.RuntimeValue)(nil) {
							value = nil
						}
						if rt, ok := value.(*ast.RuntimeValue); ok {
							switch rt.Kind {
							case ast.RuntimeValueKindObject, ast.RuntimeValueKindDict:
								fnArgValues = append(fnArgValues, reflect.ValueOf(rt))
							default:
								fnArgValues = append(fnArgValues, reflect.ValueOf(rt.Value))
							}
							continue
						}

						if rt, ok := value.(*Result); ok {
							valuesTemp := []*ast.RuntimeValue{}
							for _, result := range rt.Values {
								valuesTemp = append(valuesTemp, result.Value.(*ast.RuntimeValue))
							}
							value = valuesTemp
						}

						if rt, ok := value.([]*ast.RuntimeValue); ok {
							if isVariadicInsertion {
								for _, v := range rt {
									fnArgValues = append(fnArgValues, reflect.ValueOf(v.Value))
								}
							} else {
								fnArgValues = append(fnArgValues, reflect.ValueOf(rt))
							}
							continue
						}

						log.Warnf("Argument %d passed to runtime bound function %s::%s is not a runtime value", i, objName, methodDecl.Name)
					} else {
						log.Warnf("Argument %d passed to runtime bound function %s::%s is not a basic type", i, objName, methodDecl.Name)
					}
				}

				// args[0] is the environment
				// Which isn't exposed to the declaration/call args
				fnArgValues = append([]reflect.Value{reflect.ValueOf(args[0])}, fnArgValues...)

				returnVal := mVal.Call(fnArgValues)
				if len(returnVal) > 0 {
					return returnVal[0].Interface()
				}
				return nil
			}
		}

		objDecl.Methods[methodDecl.Name] = methodDecl
	}

	Registry.SetObject(objDecl, true)

	// log.Debugf("Registered runtime bound type %s", objName)

	for _, declaration := range objDecl.Methods {
		Registry.SetFunction(declaration)
		// log.Debugf("Registered runtime bound method %s::%s", objName, declaration.Name)
	}
}

func formatLogRuntimeValue(value any) any {
	switch rv := value.(type) {

	case *ast.RuntimeValue:
		if rv == nil {
			return nil
		}
		v, err := MarshalRuntimeValue(rv)
		if err != nil {
			log.Fatalf("Cannot marshal runtime value: %v", err)
		}
		return string(v)

	case []*ast.RuntimeValue:
		vals := make([]any, len(rv))
		for i, v := range rv {
			vv, err := MarshalRuntimeValue(v)
			if err != nil {
				log.Fatalf("Cannot marshal runtime value: %v", err)
			}
			vals[i] = string(vv)
		}
		return vals
	// case map[string]*ast.RuntimeValue:
	// 	v, err := MarshalRuntimeValue(env, rv)
	// 	if err != nil {
	// 		log.Fatalf("Cannot marshal runtime value: %v", err)
	// 	}
	// 	return string(v)

	default:
		return value
	}
}

var runtimeBindingFuncs = map[string]func(args ...any){}

func RegisterRuntimeFunctions(env *Environment) {

	bindStaticRuntimeType(new(RT_fmt))
	bindStaticRuntimeType(new(RT_error))
	bindStaticRuntimeType(new(RT_string))

	/*env.DefineCustomFunctionWithReceiver(
		ast.NewTypedIdentifierCustom("fmt", "fmt"),
		"printf",
		func(args ...interface{}) interface{} {
			if len(args) == 0 {
				panic("printf: no arguments")
			}

			fmtStr := args[0].(*ast.RuntimeValue).Value.(string)
			if len(args) == 1 {
				fmt.Print(fmtStr)
				return nil
			}

			var printArgs func(args ...any) []any
			printArgs = func(args ...any) []any {
				var argValues []any

				for _, arg := range args {
					switch rv := arg.(type) {
					case *ast.RuntimeValue:
						if rv == nil {
							argValues = append(argValues, nil)
							continue
						}
						v, err := MarshalRuntimeValue(env, rv)
						if err != nil {
							log.Fatalf("Cannot marshal runtime value: %v", err)
						}
						argValues = append(argValues, string(v))
					case []*ast.RuntimeValue:
						vals := make([]any, len(rv))
						for i, v := range rv {
							// if _, ok := v.Value.(*ast.RuntimeValue); !ok {
							// 	vals[i] = append(argValues, v.Value)
							// 	continue
							// }

							vv, err := MarshalRuntimeValue(env, v)
							if err != nil {
								log.Fatalf("Cannot marshal runtime value: %v", err)
							}
							vals[i] = string(vv)
						}
						argValues = append(argValues, vals)
					default:
						argValues = append(argValues, arg)
					}
				}

				return argValues
			}

			fmt.Printf(fmtStr, printArgs(args[1:]...)...)

			return nil
		},
	)*/

}

func init() {
	runtimeBindingFuncs["http"] = func(args ...any) {
		// objDecl := args[0].(*ast.ObjectDeclaration)
		fnDecl := args[1].(*ast.FunctionDeclaration)

		switch fnDecl.Name {

		case "get":
			fnDecl.CustomFuncCb = func(args ...any) any { return handleRequest("GET", args...) }
		case "post":
			fnDecl.CustomFuncCb = func(args ...any) any { return handleRequest("POST", args...) }
		case "put":
			fnDecl.CustomFuncCb = func(args ...any) any { return handleRequest("PUT", args...) }
		case "patch":
			fnDecl.CustomFuncCb = func(args ...any) any { return handleRequest("PATCH", args...) }
		case "delete":
			fnDecl.CustomFuncCb = func(args ...any) any { return handleRequest("DELETE", args...) }
		default:
			log.Warnf("http: unsupported method %s", fnDecl.Name)
		}
	}
}

func handleRequest(method string, args ...any) any {
	t := utilities.NewTimer("http." + method)
	defer t.StopAndLog()

	_ = args[0].(*Environment)
	if len(args) < 2 {
		log.Warnf("http.get: no arguments")
		return nil
	}
	url, ok := args[1].(*ast.RuntimeValue)
	if !ok {
		log.Warnf("http.get: argument 1 is not a runtime value")
		return nil
	}

	callExpr := ast.FindFirstParentOfType[*ast.CallExpression](url.OriginalNode)
	if callExpr == nil {
		log.Warnf("http.get: cannot find outer call expression")
		return nil
	}

	var requestOptions *ast.RuntimeValue
	if len(args) > 2 {
		requestOptions, ok = args[2].(*ast.RuntimeValue)
		if !ok {
			log.Warnf("http.get: argument 2 is not a runtime value")
			return nil
		}

		if !requestOptions.IsDictKind() {
			log.Warnf("http.get: argument 2 is not a dict")
			return nil
		}
	}

	var clientOptions []gohttpclient.ClientOption
	var options []gohttpclient.Option

	if requestOptions != nil {
		optsDict := ast.RuntimeValueAs[map[string]*ast.RuntimeValue](requestOptions)

		if headers, ok := optsDict["headers"]; ok && headers.IsDictKind() {
			headersDict := ast.RuntimeValueAs[map[string]*ast.RuntimeValue](headers)
			for k, v := range headersDict {
				if v.IsStringKind() {
					options = append(options, gohttpclient.WithHeader(k, v.Value.(string)))
				}
			}
		}

		if timeout, ok := optsDict["timeout"]; ok && timeout.IsIntegerKind() {
			clientOptions = append(clientOptions, gohttpclient.WithTimeout(time.Duration(timeout.Value.(int))*time.Millisecond))
		}

		if body, ok := optsDict["body"]; ok {
			bodyBytes, err := MarshalRuntimeValue(body)
			if err != nil {
				log.Warnf("http.get: cannot marshal body: %v", err)
				return nil
			}
			options = append(options, gohttpclient.WithBody(bodyBytes))
		}

	}

	ctx := context.Background()
	client := gohttpclient.New("", clientOptions...)

	var response *gohttpclient.Response
	var err error
	switch method {
	case "GET":
		response, err = client.Get(ctx, url.Value.(string), options...)
	case "POST":
		response, err = client.Post(ctx, url.Value.(string), options...)
	case "PUT":
		response, err = client.Put(ctx, url.Value.(string), options...)
	case "PATCH":
		response, err = client.Patch(ctx, url.Value.(string), options...)
	case "DELETE":
		response, err = client.Delete(ctx, url.Value.(string), options...)
	default:
		log.Warnf("http.get: unsupported method %s", method)
		return nil
	}
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	if response == nil {
		log.Warnf("http.get: response is nil")
		return nil
	}

	defer response.Get().Body.Close()

	assignment, ok := callExpr.GetParent().(*ast.AssignmentStatement)
	if !ok {
		log.Warnf("http.get: cannot find assignment statement, to automatically handle unmarshal to type")
		return nil
	}
	val, err := UnmarshalRuntimeValueAs(response.Body(), assignment.Type)
	if err != nil {
		log.Warnf("http.get: cannot unmarshal response body: %v", err)
		return nil
	}

	return val
}
