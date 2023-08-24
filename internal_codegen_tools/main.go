package main

import (
	"github.com/jessevdk/go-flags"
)

type CliOptions struct {
	Tool string `long:"tool" choice:"visitor-generator"`

	DebugData            bool `long:"debug" description:"Prints Debug data"`
	DebugDataDumpStructs bool `long:"debug-dump-structs" description:"Prints structs when printing Debug data"`
	DebugBuiltData       bool `long:"debug-built-data" description:"Prints built data when printing Debug data"`

	DryRun bool `short:"d" long:"dry" description:"Dry run, don't write any files, outputs to cli"`
}

var opts = &CliOptions{}

func main() {
	if _, err := flags.Parse(opts); err != nil {
		panic(err)
	}

	tools := map[string]ICliTool{
		"visitor-generator": NewAstVisitorGenerator(),
	}

	tool, ok := tools[opts.Tool]
	if !ok {
		panic("invalid tool: " + opts.Tool)
	}

	tool.SetOptions(opts)
	tool.RunTool()

}
