package main

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type enumValData struct {
	Name  string
	Value int
}

func main() {
	var enumVals []enumValData
	filterFunc := func(info os.FileInfo) bool {
		name := info.Name()
		if info.IsDir() {
			return false
		}
		if filepath.Ext(name) != ".go" {
			return false
		}
		if strings.HasSuffix(name, "_test.go") {
			return false
		}
		return true
	}

	pkgs, err := parser.ParseDir(token.NewFileSet(), "../", filterFunc, 0)

	if err != nil {
		panic("Failed to parse grammar package: " + err.Error())
	}
	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
			// Now we need to extract the parser rules from the package	// Example:
			// const (
			//   SimpleLangParserRULE_program                    = 0
			//   SimpleLangParserRULE_typedIdentifier            = 1
			// )

			ast.Inspect(file, func(n ast.Node) bool {
				// Check if the node is a GenDecl (General Declaration)
				decl, ok := n.(*ast.GenDecl)
				if !ok || decl.Tok != token.CONST {
					// We're only interested in constant declarations
					return true
				}
				for _, spec := range decl.Specs {
					// Extracting value specification (constant)
					valSpec, ok := spec.(*ast.ValueSpec)
					if !ok {
						continue
					}
					for _, name := range valSpec.Names {
						if !strings.HasPrefix(name.Name, "SimpleLangParserRULE_") {
							continue
						}
						val := valSpec.Values[0]
						if val == nil {
							continue
						}

						litVal := val.(*ast.BasicLit)
						if litVal == nil {
							continue
						}

						if litVal.Kind == token.INT {
							i, err := strconv.Atoi(litVal.Value)
							if err != nil {
								panic(err)
							}

							// ensure that the enum value is not already in the list
							exists := false
							for _, v := range enumVals {
								if v.Name == name.Name {
									exists = true
									break
								}
							}
							if exists {
								continue
							}

							newName := name.Name
							if strings.HasPrefix(newName, "SimpleLangParserRULE_") {
								newName = newName[len("SimpleLangParserRULE_"):]
								newName = strings.ToUpper(newName[:1]) + newName[1:]
								newName = "ParserRule" + newName
							}

							enumVals = append(enumVals, enumValData{
								Name:  newName,
								Value: i,
							})
						}
					}
				}
				return true
			})
		}
	}

	// Now we need to generate the enum

	fileData := `package grammar

type ParserRule int

const (
`
	for _, val := range enumVals {
		fileData += fmt.Sprintf("\t%s ParserRule = %d\n", val.Name, val.Value)
	}
	fileData += `)`

	fileData += `
func (self ParserRule) String() string {
	switch self {
`
	for _, v := range enumVals {
		fileData += fmt.Sprintf("\tcase %s: return \"%s\"\n", v.Name, v.Name)
	}
	fileData += `	}
	return ""
}
`

	// Now we format the file
	fileDataBytes, err := format.Source([]byte(fileData))
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("../parser_rules_enum.go", fileDataBytes, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Successfully generated parser_rules_enum.go\n")
}
