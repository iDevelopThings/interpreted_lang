package main

import (
	"flag"
	"os"
	"runtime/pprof"

	"github.com/charmbracelet/log"

	"arc/interpreter"
	"arc/lsp"
)

var env = interpreter.NewEnvironment()

var cpuprofile = flag.Bool("cpuprofile", false, "write cpu profile to file")
var memprofile = flag.Bool("memprofile", false, "write memory profile to this file")
var runLsp = flag.Bool("lsp", false, "run the lsp only")
var lspProtocol = flag.String("lsp-protocol", "stdio", "run the lsp using the protocol")

func main() {
	flag.Parse()

	if *runLsp {
		lsp.Run(*lspProtocol)
		return
	}

	if *cpuprofile {
		f, err := os.Create("cpu.prof")
		if err != nil {
			log.Fatal(err)
		}
		err = pprof.StartCPUProfile(f)
		if err != nil {
			log.Fatalf("could not start cpu profile: %v", err)
			return
		}
		defer pprof.StopCPUProfile()
	}

	engine := interpreter.Engine
	// engine.LoadScript("test_data/http_basic_test.arc")
	// engine.LoadScript("test_data/http.arc")
	// engine.LoadScript("test_data/type_checking.arc")
	engine.LoadScript("test_data/testing.arc")
	// engine.LoadScript("test_data/imports.arc")
	// engine.LoadScript("test_data/loops.arc")
	// engine.LoadScript("test_data/dictionaries.arc")
	// engine.LoadScript("test_data/input.arc")
	engine.Run()

	if *memprofile {
		f, err := os.Create("mem.mprof")
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
	}

}
