package config

import (
	"encoding/json"
	"errors"
	"strings"
)

type OutputFormat string

const (
	OutputFormatJson OutputFormat = "json"
	OutputFormatText OutputFormat = "text"
)

type CliArgsConfig struct {
	Verbose []bool `short:"v" long:"verbose" description:"Show verbose debug information"`

	// An override to use for the working dir
	// instead of os.Getwd()
	WorkingDirectory string `short:"d" long:"dir" description:"The working directory to use"`

	CpuProfile bool `long:"cpuprofile" description:"write cpu profile to file"`
	MemProfile bool `long:"memprofile" description:"write memory profile to this file"`

	RunLsp      bool   `long:"lsp" description:"run the lsp only"`
	LspProtocol string `long:"lsp-protocol" description:"run the lsp using the protocol"`

	PrintAst     bool         `long:"print-ast" description:"print the ast to stderr"`
	StdinMode    bool         `short:"i" long:"stdin" description:"Only accept input from stdin"`
	LintingMode  bool         `short:"l" long:"lint" description:"run the linter only"`
	OutputFormat OutputFormat `short:"f" long:"format" description:"the output format to use for the linter" choice:"json" choice:"text" default:"text"`

	File string
}

type EnvProxiedValue[T any] struct {
	Value   T
	EnvName string
}

func (self *EnvProxiedValue[T]) UnmarshalJSON(data []byte) error {
	rawValue := string(data)

	// We can pass a raw value, for ex, the specified port:
	// `132`
	// Or we can specific what the env variable name is:
	// `env:PORT`

	if len(rawValue) == 0 {
		return errors.New("EnvProxiedValue cannot be empty")
	}

	rawValue = strings.TrimSpace(rawValue)
	if rawValue[0] == '"' && rawValue[len(rawValue)-1] == '"' {
		rawValue = rawValue[1 : len(rawValue)-1]
	}

	if strings.HasPrefix(rawValue, "env:") {
		self.EnvName = rawValue[4:]
		return nil
	}

	val := (*T)(nil)
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}

	self.Value = *val

	return nil
}

type ProjectConfiguration struct {
	// Not required, but nice to have
	ProjectName string `json:"project_name"`

	HttpServer *HttpServerConfiguration `json:"http_server"`
}

type HttpServerConfiguration struct {
	Port    *EnvProxiedValue[int]    `json:"port"`
	Address *EnvProxiedValue[string] `json:"address"`

	// The max memory size in bytes for the uploaded form files/data
	// Default: 10<<20 (10 MB)
	FormMaxMemory int64 `json:"form_max_memory"`
}
