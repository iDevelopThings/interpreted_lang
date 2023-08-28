package config

import (
	"encoding/json"
	"errors"
	"strings"
)

type OutputFormat string

const (
	OutputFormatJson       OutputFormat = "json"
	OutputFormatJsonIndent OutputFormat = "json-indent"
	OutputFormatText       OutputFormat = "text"
)

type LogOutputMode string

const (
	LogOutputModeStdErr LogOutputMode = "stderr"
	LogOutputModeStdOut LogOutputMode = "stdout"
)

type CliArgsConfig struct {
	Verbose []bool `short:"v" long:"verbose" description:"Show verbose debug information"`

	// An override to use for the working dir
	// instead of os.Getwd()
	WorkingDirectory string `short:"d" long:"dir" description:"The working directory to use"`

	CpuProfile bool `long:"cpuprofile" description:"write cpu profile to file"`
	MemProfile bool `long:"memprofile" description:"write memory profile to this file"`

	PrintAst bool `long:"print-ast" description:"print the ast to stderr"`

	StdinMode bool `short:"i" long:"stdin" description:"Only accept input from stdin"`

	LintingMode bool `short:"l" long:"lint" description:"run the linter only"`

	OutputFormat OutputFormat `short:"f" long:"format" description:"the output format to use for the linter" choice:"json" choice:"text" choice:"json-indent" default:"text"`

	// This is mainly for use when we're linting
	// If we're linting, we only want to receive typecheck/parse errors as json via stdout
	// Everything else should be via stderr
	LogOutputMode LogOutputMode `long:"log-output" description:"the output mode to use for the logger" choice:"stderr" choice:"stdout" default:"stderr"`

	Extra struct {
		File string `positional-arg-name:"file" description:"The arc source to run"`
	} `positional-args:"yes"`
}

type ProjectConfiguration struct {
	// Not required, but nice to have
	ProjectName string `json:"project_name"`

	HttpServer *HttpServerConfiguration `json:"http_server"`
}

func createDefaultProjectConfig() *ProjectConfiguration {
	return &ProjectConfiguration{
		ProjectName: "",
		HttpServer: &HttpServerConfiguration{
			Port:          &EnvProxiedValue[int]{Value: 8080},
			Address:       &EnvProxiedValue[string]{Value: "localhost"},
			FormMaxMemory: 10 << 20, // 10 MB

			ReadHeaderTimeout: &EnvProxiedValue[int64]{Value: 5000},
			WriteTimeout:      &EnvProxiedValue[int64]{Value: 5000},
		},
	}
}

type HttpServerConfiguration struct {
	Port    *EnvProxiedValue[int]    `json:"port"`
	Address *EnvProxiedValue[string] `json:"address"`

	// The max memory size in bytes for the uploaded form files/data
	// Default: 10<<20 (10 MB)
	FormMaxMemory int64 `json:"form_max_memory"`

	// Values to use with `http.Server` options
	// Timeouts are in ms
	ReadHeaderTimeout *EnvProxiedValue[int64] `json:"read_header_timeout_ms"`
	WriteTimeout      *EnvProxiedValue[int64] `json:"write_timeout_ms"`
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
