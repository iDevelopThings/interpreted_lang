package main

import (
	"os"
	"os/signal"
	"runtime/pprof"
	"syscall"

	"arc/log"

	"arc/interpreter"
	"arc/interpreter/config"
	"arc/interpreter/errors"
	"arc/utilities"
)

func main() {

	runTimer := utilities.NewTimer("Process Execution")
	defer runTimer.StopAndLog()

	cliConf := config.PrepareConfiguration()
	config.LoadProjectConfiguration()

	if cliConf.CpuProfile {
		runCpuProfiling()
		defer pprof.StopCPUProfile()
	}
	if cliConf.MemProfile {
		defer memProfile()
	}

	// engine := interpreter.Engine
	// engine.LoadSource(cliConf.File)
	// engine.LoadSource("test_data/http.arc")
	// engine.LoadSource("test_data/http_basic_test.arc")
	// engine.LoadSource("test_data/type_checking.arc")
	// engine.LoadSource("test_data/testing.arc")
	// engine.LoadSource("test_data/errors_pls.arc")
	// engine.LoadSource("test_data/enums.arc")
	// engine.LoadSource("test_data/imports.arc")
	// engine.LoadSource("test_data/loops.arc")
	// engine.LoadSource("test_data/dictionaries.arc")
	// engine.LoadSource("test_data/input.arc")
	// engine.Run()

	// read all contents from stdin

	runEngineAndScript(cliConf)

}

func runEngineAndScript(conf *config.CliArgsConfig) {
	errors.SetFormat(conf.OutputFormat)

	engine := interpreter.Engine
	engine.Load()

	if conf.LintingMode {
		engine.Lint()
	} else {
		engine.Run()
	}

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // subscribe to system signals
	onKill := func(c chan os.Signal) {
		<-c
		defer pprof.StopCPUProfile()
		defer os.Exit(0)
	}

	// try to handle os interrupt(signal terminated)
	go onKill(c)
}

func runCpuProfiling() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	err = pprof.StartCPUProfile(f)
	if err != nil {
		log.Fatalf("could not start cpu profile: %v", err)
	}
}

func memProfile() {
	f, err := os.Create("mem.mprof")
	if err != nil {
		log.Fatal(err)
	}
	err = pprof.WriteHeapProfile(f)
	if err != nil {
		log.Fatalf("could not write heap profile: %v", err)
	}
	err = f.Close()
	if err != nil {
		log.Fatalf("could not close file: %v", err)
	}
}
