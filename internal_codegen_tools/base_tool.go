package main

import (
	"os"
)

type ICliTool interface {
	SetOptions(o *CliOptions)
	RunTool()
}

type CliTool struct {
	GoFile    string
	GoPackage string

	DryRun               bool
	DebugData            bool
	DebugDataDumpStructs bool
	DebugBuiltData       bool
}

func (t *CliTool) SetOptions(o *CliOptions) {
	t.DryRun = o.DryRun
	t.DebugData = o.DebugData
	t.DebugDataDumpStructs = o.DebugDataDumpStructs
	t.DebugBuiltData = o.DebugBuiltData
}

func NewBaseCliTool() *CliTool {
	t := &CliTool{}
	t.GoFile = os.Getenv("GOFILE")
	t.GoPackage = os.Getenv("GOPACKAGE")

	return t
}
