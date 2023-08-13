package interpreter

import (
	"fmt"

	"github.com/charmbracelet/log"

	"interpreted_lang/ast"
)

func RegisterRuntimeFunctions(env *Environment) {
	// for _, kind := range ast.AllLiteralKinds {
	// 	env.SetObject(&ast.ObjectDeclaration{Name: string(kind)})
	// }

	env.DefineCustomFunctionWithReceiver(
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
	)
}
