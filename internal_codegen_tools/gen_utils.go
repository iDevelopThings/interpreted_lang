package main

import (
	"bytes"
	"go/format"
	"os"
	"path/filepath"
	"strings"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
)

func formatAndWriteFile(f *dst.File, filename string, dryRun bool) string {
	// fset := token.NewFileSet()
	var buf bytes.Buffer

	buf.WriteString("// Code generated by arc parser tooling. DO NOT EDIT.\n")
	fset, file, err := decorator.RestoreFile(f)
	if err != nil {
		log.Fatalf("decorator.RestoreFile: %s", err)
	}
	// formatted, err := format.Source(buf.Bytes())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// buf = *bytes.NewBuffer(formatted)

	if err := format.Node(&buf, fset, file); err != nil {
		log.Fatal(err)
	}

	if dryRun {
		return buf.String()
	}

	if err := os.WriteFile(filename, buf.Bytes(), 0644); err != nil {
		log.Fatal(err)
	}

	cwd, _ := os.Getwd()
	cwd = filepath.Join(cwd, "..")

	log.Debugf("Wrote file %s", strings.Replace(filename, cwd, ".", 1))

	return buf.String()
}
