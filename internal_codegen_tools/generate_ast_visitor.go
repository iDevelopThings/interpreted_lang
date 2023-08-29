package main

import (
	"bytes"
	"embed"
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"text/template"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/davecgh/go-spew/spew"
)

type AstVisitorGenerator struct {
	*CliTool

	File    *dst.File
	Locator *AstNodeLocator

	VisitorStructs map[string]*visitorStructData
	StructsData    []*visitorStructData
	VisitorMethods []*VisitorMethod
}

func NewAstVisitorGenerator() ICliTool {
	tool := &AstVisitorGenerator{
		CliTool: NewBaseCliTool(),
	}

	return tool
}

var (
	//go:embed templates/*.go.tmpl
	templates embed.FS
)

func (self *AstVisitorGenerator) RunTool() {
	self.File = &dst.File{Name: dst.NewIdent("ast")}
	self.VisitorStructs = map[string]*visitorStructData{}

	self.Locator = NewAstNodeLocator(self.GoFile, nil)
	self.Locator.ExcludedFileSuffixes = []string{".gen.go"}
	self.Locator.ExcludedInterfaceNames = []string{
		"Node",
		"NodeVisitor",
		"Statement",
		"TopLevelStatement",
		"Expr",
		"Declaration",
		"ObjectFieldGetter",
		"Type",
	}
	self.Locator.ExcludedStructNames = []string{
		"AstNode",
		"RuntimeValue",
		"BasicType",
		"TokenRange",
		"ParserRuleRange",
		"HttpRequestObject",
		"NodeVisitorAdapter",
	}
	self.Locator.Locate()

	if self.DebugData {
		spew.Config.Indent = "  "
		spew.Config.MaxDepth = 3
		spew.Config.DisablePointerAddresses = true

		log.Debugf("Located structs:")
		for name, astStruct := range self.Locator.Structs {
			s := fmt.Sprintf("  %s", name)
			if self.DebugDataDumpStructs {
				s += fmt.Sprintf(":\n%s", spew.Sdump(astStruct))
			}
			log.Debugf(s)
		}

		println()
		println()
		log.Debugf("Located interfaces:")
		for name, astInterface := range self.Locator.Interfaces {
			s := fmt.Sprintf("  %s", name)
			if self.DebugDataDumpStructs {
				s += fmt.Sprintf(":\n%s", spew.Sdump(astInterface))
			}
			log.Debugf(s)
		}

		os.Exit(0)
	}

	self.buildData()
	self.Generate()
}

func (self *AstVisitorGenerator) buildData() {
	var structsData []*visitorStructData
	var visitorMethods []*VisitorMethod

	for s, astStruct := range self.Locator.Structs {
		visitorMethods = append(visitorMethods, &VisitorMethod{
			TypeName:    astStruct.Name,
			IsArray:     false,
			IsPtr:       true,
			IsInterface: false,
		})

		visitorStructDat := &visitorStructData{
			AstStruct: astStruct,
			VisitArgs: []visitorArg{},
		}
		for _, field := range astStruct.st.Fields.List {
			if field.Tag != nil {
				tag := reflect.StructTag(field.Tag.Value[1 : len(field.Tag.Value)-1])

				tagValue, ok := tag.Lookup("visitor")
				if ok && tagValue == "-" {
					continue
				}
			}

			typeInfo := resolveTypeInfo(field.Type)

			fieldName := ""
			if len(field.Names) > 0 {
				fieldName = field.Names[0].Name
			} else {
				fieldName = typeInfo.Name
			}

			if fieldName == "AstNode" {
				continue
			}

			if typeInfo.IsInterface == false {
				typeInfo.IsInterface = self.Locator.IsInterface(typeInfo.Name)
			}
			visitorStructDat.VisitArgs = append(visitorStructDat.VisitArgs, visitorArg{
				StructKey:       fieldName,
				Type:            typeInfo.Name,
				IsArray:         typeInfo.IsArray,
				IsMap:           typeInfo.IsMap,
				IsPtr:           typeInfo.IsPtr,
				IsInterfaceType: typeInfo.IsInterface,
				IsStructType:    typeInfo.IsStruct,
			})
		}

		structsData = append(structsData, visitorStructDat)
		self.VisitorStructs[s] = visitorStructDat
	}

	sort.Slice(structsData, func(i, j int) bool {
		return typeCategory(structsData[i].Name) < typeCategory(structsData[j].Name)
	})
	sort.Slice(visitorMethods, func(i, j int) bool {
		return visitorMethodTypeCategory(visitorMethods[i]) < visitorMethodTypeCategory(visitorMethods[j])
	})

	self.StructsData = structsData
	self.VisitorMethods = visitorMethods

	if self.DebugBuiltData {
		spew.Config.Indent = "  "
		spew.Config.MaxDepth = 4
		spew.Config.DisablePointerAddresses = true

		if self.DebugDataDumpStructsFields {
			tmpFile := &dst.File{Name: dst.NewIdent("ast")}
			for _, s := range self.Locator.StructDeclarations {

				tmpFile.Decls = append(tmpFile.Decls, s)
				// buf := new(bytes.Buffer)
				// if err := format.Node(buf, token.NewFileSet(), s.AstStruct.st); err != nil {
				// 	log.Fatalf("format.Node: %s", err)
				// }
				// fmt.Printf("%s:\n%s\n", s.Name, buf.String())
			}

			decorator.Print(tmpFile)

			os.Exit(0)
		}

		log.Debugf("StructsData: ")
		spew.Dump(self.StructsData)
		log.Debugf("VisitorMethods: ")
		spew.Dump(self.VisitorMethods)

		os.Exit(0)
	}
}

func (self *AstVisitorGenerator) makeGoAstFromTemplate() {
	t, err := template.ParseFS(templates, "templates/walk_func.go.tmpl")
	if err != nil {
		log.Fatalf("template.ParseFS: %s", err)
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, self); err != nil {
		log.Fatalf("t.Execute: %s", err)
	}

	writer := strings.Builder{}
	writer.WriteString(buf.String())

	fset := token.NewFileSet()
	f, err := decorator.ParseFile(fset, "", writer.String(), parser.ParseComments)
	if err != nil {
		log.Debugf("template: %s", writer.String())
		log.Fatalf("parser.ParseFile: %s", err)
	}

	self.File.Decls = append(self.File.Decls, f.Decls...)
}

func (self *AstVisitorGenerator) Generate() {
	self.makeGoAstFromTemplate()

	d := filepath.Dir(self.GoFile)

	fileOut := formatAndWriteFile(
		self.File,
		filepath.Join(d, "..", "ast", "node_visitor_accept_functions.gen.go"),
		self.DryRun,
	)

	if self.DryRun {
		log.Debugf("Output:\n%s\n", fileOut)
	}

}

func visitorMethodTypeCategory(method *VisitorMethod) int {

	if method.TypeName == "Program" {
		return 0
	}
	if strings.HasSuffix(method.TypeName, "TopLevelStatement") || strings.HasSuffix(method.TypeName, "Declaration") {
		if method.IsInterface {
			return 1
		}
		return 2
	}
	if strings.HasSuffix(method.TypeName, "Statement") {
		if method.IsInterface {
			return 3
		}
		return 4
	}
	if strings.HasSuffix(method.TypeName, "Expr") {
		if method.IsInterface {
			return 5
		}
		return 6
	}
	if strings.HasSuffix(method.TypeName, "Literal") {
		if method.IsInterface {
			return 7
		}
		return 8
	}

	return 6
}

func typeCategory(name string) int {
	if name == "" {
		return 0
	}
	if name == "Program" {
		return 0
	}
	if strings.HasSuffix(name, "Declaration") {
		return 0
	}
	if strings.HasSuffix(name, "Statement") {
		return 1
	}
	if strings.HasSuffix(name, "Expr") {
		return 2
	}
	if strings.HasSuffix(name, "Literal") {
		return 3
	}

	return 4
}
