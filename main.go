package main

import (
	"flag"
	"os"
	"runtime/pprof"

	"github.com/charmbracelet/log"

	"interpreted_lang/interpreter"
)

var env = interpreter.NewEnvironment()

var cpuprofile = flag.Bool("cpuprofile", false, "write cpu profile to file")
var memprofile = flag.Bool("memprofile", false, "write memory profile to this file")

func main() {
	flag.Parse()
	if *cpuprofile {
		f, err := os.Create("cpu.prof")
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	engine := interpreter.Engine
	// engine.LoadScript("test_data/http_basic_test.sl")
	// engine.LoadScript("test_data/http.sl")
	// engine.LoadScript("test_data/type_checking.sl")
	engine.LoadScript("test_data/imports.sl")
	// engine.LoadScript("test_data/loops.sl")
	// engine.LoadScript("test_data/dictionaries.sl")
	// engine.LoadScript("test_data/input.sl")
	engine.Run()

	if *memprofile {
		f, err := os.Create("mem.mprof")
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
	}

	// startedAt := time.Now()
	//
	// src, err := os.ReadFile("test_data/http.sl")
	// if err != nil {
	// 	panic(err)
	// }
	// is := antlr.NewInputStream(string(src))
	// lexer := grammar.NewSimpleLangLexer(is)
	// stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// tree := grammar.NewSimpleLangParser(stream)
	// programTree := tree.Program()
	//
	// fmt.Println("Lex & Parse time:", time.Now().Sub(startedAt))

}
