package main

import (
	"runtime"
	"testing"
)

func Test_Main(t *testing.T) {
	runtime.SetCPUProfileRate(500)
	// runEngineAndScript("test_data/1m_funcs.arc")

}
