package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	l "github.com/charmbracelet/log"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
)

var log *l.Logger = l.NewWithOptions(os.Stderr, l.Options{
	ReportCaller:    true,
	ReportTimestamp: true,
	TimeFormat:      time.Kitchen,
	Prefix:          "AstLocator",
	Level:           l.DebugLevel,
})

type AstStruct struct {
	Name string
	st   *dst.StructType
}

type AstNodeLocator struct {
	Dir                    string
	ExcludedFileSuffixes   []string
	Structs                map[string]*AstStruct
	Interfaces             map[string]*dst.InterfaceType
	ExcludedInterfaceNames []string
	ExcludedStructNames    []string

	Checker func(typ *dst.TypeSpec) bool
}

var AllLocatedInterfaces = map[string]bool{}
var AllLocatedStructs = map[string]bool{}

func NewAstNodeLocator(dir string, checker func(typ *dst.TypeSpec) bool) *AstNodeLocator {
	inst := &AstNodeLocator{
		Dir:                  filepath.Dir(dir),
		ExcludedFileSuffixes: []string{},
		Structs:              make(map[string]*AstStruct),
		Interfaces:           make(map[string]*dst.InterfaceType),
		Checker:              checker,
	}
	if inst.Checker == nil {
		inst.Checker = func(typ *dst.TypeSpec) bool {
			return true
		}
	}

	return inst
}

func (self *AstNodeLocator) isExcluded(fileName string) bool {
	for _, suffix := range self.ExcludedFileSuffixes {
		if strings.HasSuffix(fileName, suffix) {
			return true
		}
	}
	return false
}

func (self *AstNodeLocator) Locate() {
	fset := token.NewFileSet()

	// pkgs, err := parser.ParseDir(fset, self.Dir, nil, parser.ParseComments)
	// if err != nil {
	// 	panic(err)
	// }
	/*
		// Invoke the type checker using AST as input
		typesInfo := types.Info{
			Defs: make(map[*ast.Ident]types.Object),
			Uses: make(map[*ast.Ident]types.Object),
		}
		conf := &types.Config{}

		pkgFiles := []*ast.File{}
		for _, pkg := range pkgs {
			for _, file := range pkg.Files {
				pkgFiles = append(pkgFiles, file)
			}
		}

		if _, err := conf.Check("", fset, pkgFiles, &typesInfo); err != nil {
			var list types.Error
			if errors.As(err, &list) {
				if strings.HasPrefix(list.Msg, "could not import") {
					log.Warnf("Parsing File: Could not import %s", list.Msg[16:])
				} else {
					panic(err)
				}
			} else {
				panic(err)
			}
		}

		// Create a new decorator, which will track the mapping between ast and dst nodes
		dec := decorator.NewDecorator(fset)

		for _, pkg := range pkgs {
			for fileName, f := range pkg.Files {
				if self.isExcluded(fileName) {
					continue
				}
				// Decorate the *ast.File to give us a *dst.File
				file, err := dec.DecorateFile(f)
				if err != nil {
					panic(err)
				}

				for _, decl := range file.Decls {
					genDecl, ok := decl.(*dst.GenDecl)
					if !ok {
						continue
					}
					if genDecl.Tok != token.TYPE {
						continue
					}
					for _, spec := range genDecl.Specs {
						typeSpec, ok := spec.(*dst.TypeSpec)
						if !ok {
							continue
						}
						if structType, ok := typeSpec.Type.(*dst.StructType); ok {
							if !self.Checker(typeSpec) {
								continue
							}
							if self.ExcludedStructNames != nil {
								isExcluded := false
								for _, excludedName := range self.ExcludedStructNames {
									if typeSpec.Name.Name == excludedName {
										isExcluded = true
										break
									}
								}
								if isExcluded {
									continue
								}
							}

							if _, ok := self.Structs[typeSpec.Name.Name]; ok {
								log.Fatalf("Duplicate struct name: %s\n", typeSpec.Name.Name)
							}

							self.Structs[typeSpec.Name.Name] = &AstStruct{
								Name: typeSpec.Name.Name,
								st:   structType,
							}
						} else if interfaceType, ok := typeSpec.Type.(*dst.InterfaceType); ok {
							if self.ExcludedInterfaceNames != nil {
								isExcluded := false
								for _, excludedName := range self.ExcludedInterfaceNames {
									if typeSpec.Name.Name == excludedName {
										isExcluded = true
										break
									}
								}
								if isExcluded {
									continue
								}
							}
							if _, ok := self.Interfaces[typeSpec.Name.Name]; ok {
								log.Fatalf("Duplicate interface name: %s\n", typeSpec.Name.Name)
							}

							self.Interfaces[typeSpec.Name.Name] = interfaceType
						} else {
							continue
						}
					}
				}
			}
		}*/

	pkgs, err := decorator.ParseDir(fset, self.Dir, nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("Error parsing directory: %v\n", err)
	}

	for _, pkg := range pkgs {
		for fileName, file := range pkg.Files {
			if self.isExcluded(fileName) {
				continue
			}

			for _, decl := range file.Decls {
				genDecl, ok := decl.(*dst.GenDecl)
				if !ok {
					continue
				}
				if genDecl.Tok != token.TYPE {
					continue
				}
				for _, spec := range genDecl.Specs {
					typeSpec, ok := spec.(*dst.TypeSpec)
					if !ok {
						continue
					}
					if structType, ok := typeSpec.Type.(*dst.StructType); ok {
						if !self.Checker(typeSpec) {
							continue
						}
						AllLocatedStructs[typeSpec.Name.Name] = true

						if self.ExcludedStructNames != nil {
							isExcluded := false
							for _, excludedName := range self.ExcludedStructNames {
								if typeSpec.Name.Name == excludedName {
									isExcluded = true
									break
								}
							}
							if isExcluded {
								continue
							}
						}

						if _, ok := self.Structs[typeSpec.Name.Name]; ok {
							log.Fatalf("Duplicate struct name: %s\n", typeSpec.Name.Name)
						}

						self.Structs[typeSpec.Name.Name] = &AstStruct{
							Name: typeSpec.Name.Name,
							st:   structType,
						}
					} else if interfaceType, ok := typeSpec.Type.(*dst.InterfaceType); ok {
						AllLocatedInterfaces[typeSpec.Name.Name] = true
						if self.ExcludedInterfaceNames != nil {
							isExcluded := false
							for _, excludedName := range self.ExcludedInterfaceNames {
								if typeSpec.Name.Name == excludedName {
									isExcluded = true
									break
								}
							}
							if isExcluded {
								continue
							}
						}
						if _, ok := self.Interfaces[typeSpec.Name.Name]; ok {
							log.Fatalf("Duplicate interface name: %s\n", typeSpec.Name.Name)
						}

						self.Interfaces[typeSpec.Name.Name] = interfaceType
					} else {
						continue
					}
				}
			}
		}
	}
}

func (self *AstNodeLocator) IsInterface(name string) bool {
	_, ok := self.Interfaces[name]
	return ok
}

type FindFunc struct {
	Name      string
	TypeParam string
}

type ProjectWalker struct {
	RootDir string

	FindFuncs []FindFunc
	Located   []LocatedFunc
}
type LocatedFunc struct {
	Name string
	Path string
}

func NewProjectWalker(rootDir string, find []FindFunc) *ProjectWalker {
	inst := &ProjectWalker{
		RootDir:   rootDir,
		FindFuncs: find,
		Located:   []LocatedFunc{},
	}

	inst.Locate()

	return inst
}

func (self *ProjectWalker) Locate() {
	err := filepath.Walk(self.RootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Only process Go source files
		if !info.IsDir() && filepath.Ext(path) == ".go" {
			fset := token.NewFileSet()
			astFile, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
			if err != nil {
				log.Printf("Error parsing file %s: %s", path, err)
				return nil
			}

			typesInfo := types.Info{
				Defs: make(map[*ast.Ident]types.Object),
				Uses: make(map[*ast.Ident]types.Object),
			}
			conf := &types.Config{}
			if _, err := conf.Check("", fset, []*ast.File{astFile}, &typesInfo); err != nil {
				panic(err)
			}

			// Create a new decorator, which will track the mapping between ast and dst nodes
			dec := decorator.NewDecorator(fset)

			// Decorate the *ast.File to give us a *dst.File
			f, err := dec.DecorateFile(astFile)
			if err != nil {
				panic(err)
			}

			// Find all references to the function
			dst.Inspect(f, func(n dst.Node) bool {
				callExpr, ok := n.(*dst.CallExpr)
				if !ok {
					return true
				}

				ident, ok := callExpr.Fun.(*dst.Ident)
				if !ok {
					if idxExpr, ok := callExpr.Fun.(*dst.IndexExpr); ok {
						if x, ok := idxExpr.X.(*dst.SelectorExpr); ok {
							// ident, ok = selExpr.Sel
							// if !ok {
							// 	return true

							_, ok := x.X.(*dst.Ident)
							if !ok {
								return true
							}

							ident = x.Sel

						} else {
							return true
						}
					} else {
						return true
					}
				}

				if self.FindFuncs != nil {
					for _, f := range self.FindFuncs {
						if ident.Name == f.Name {
							if f.TypeParam != "" && !typeParameterPresent(callExpr.Args, f.TypeParam) {
								continue
							}
							self.Located = append(self.Located, LocatedFunc{
								Name: ident.Name,
								Path: path,
							})
							fmt.Printf("Function '%s' found in file %s\n", f.Name, path)
						}
					}
				}

				return true
			})
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

func typeParameterPresent(args []dst.Expr, typeParameter string) bool {
	for _, arg := range args {
		argType := reflect.TypeOf(arg)
		if strings.Contains(argType.String(), typeParameter) {
			return true
		}
	}
	return false
}
