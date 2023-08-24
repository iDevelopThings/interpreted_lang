package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/dave/dst"
)

type TypeInfo struct {
	IsPtr       bool
	IsArray     bool
	IsInterface bool
	IsStruct    bool
	Name        string
	IsMap       bool
}

func resolveTypeInfo(expr dst.Expr) *TypeInfo {
	typeName, isPtr, isArray, tmpIsMap, isInterface, isStruct := getTypeInfo(expr, "")

	isArray = strings.HasPrefix(reflect.TypeOf(expr).String(), "*dst.ArrayType")

	return &TypeInfo{
		IsPtr:       isPtr,
		IsArray:     isArray,
		IsInterface: isInterface,
		IsStruct:    isStruct,
		IsMap:       tmpIsMap,
		Name:        typeName,
	}
}

func getTypeInfo(expr dst.Expr, currentName string) (typeName string, isPtr, isArray, isMap, isInterface, isStruct bool) {
	tmpTypeName := ""
	tmpIsPtr := false
	tmpIsArray := false
	tmpIsMap := false
	tmpIsInterface := false
	tmpIsStruct := false

	switch typ := expr.(type) {
	case *dst.StarExpr:
		tmpTypeName, tmpIsPtr, tmpIsArray, tmpIsMap, tmpIsInterface, tmpIsStruct = getTypeInfo(typ.X, currentName)

		isPtr = true
	case *dst.Ident:
		tmpTypeName = typ.Name
		// Determine if the type is an interface type
		if typ.Obj != nil {
			if typ.Obj.Kind == dst.Typ && typ.Obj.Decl != nil {
				if ts, ok := typ.Obj.Decl.(*dst.TypeSpec); ok {
					if _, ok := ts.Type.(*dst.InterfaceType); ok {
						tmpIsInterface = true
					} else if _, ok := ts.Type.(*dst.StructType); ok {
						tmpIsStruct = true
					}
				}
			}
		} else {
			if _, ok := AllLocatedStructs[typ.Name]; ok {
				tmpIsStruct = true
			} else if _, ok := AllLocatedInterfaces[typ.Name]; ok {
				tmpIsInterface = true
			}
		}
	case *dst.MapType:
		tmpTypeName, tmpIsPtr, tmpIsArray, tmpIsMap, tmpIsInterface, tmpIsStruct = getTypeInfo(typ.Value, currentName)
		tmpIsMap = true
	case *dst.ArrayType:
		tmpTypeName, tmpIsPtr, tmpIsArray, tmpIsMap, tmpIsInterface, tmpIsStruct = getTypeInfo(typ.Elt, currentName)

		// switch typ2 := typ.Elt.(type) {
		// case *dst.Ident:
		// 	typeName = typ2.Name
		// case *dst.StarExpr:
		// 	isPtr = true
		// 	switch typ3 := typ2.X.(type) {
		// 	case *dst.Ident:
		// 		typeName = typ3.Name
		// 	case *dst.SelectorExpr:
		// 		typeName = fmt.Sprintf("%s.%s", typ3.X.(*dst.Ident).Name, typ3.Sel.Name)
		// 	}
		// }
	}

	if tmpTypeName == "" {
		typeName = currentName
	} else {
		// If currentName isn't empty & tmpTypeName isn't, then it's a nested type (e.g. a.b.c)
		if currentName != "" {
			typeName = fmt.Sprintf("%s.%s", currentName, tmpTypeName)
		} else {
			typeName = tmpTypeName
		}
	}

	isPtr = tmpIsPtr
	isArray = tmpIsArray
	isMap = tmpIsMap
	isInterface = tmpIsInterface
	isStruct = tmpIsStruct

	// We need to check if the type is an interface type

	return
}
