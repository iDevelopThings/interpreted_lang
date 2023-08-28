package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"

	"github.com/charmbracelet/log"
	"github.com/jessevdk/go-flags"
)

var CliConfig *CliArgsConfig = &CliArgsConfig{}
var ProjectConfig *ProjectConfiguration = createDefaultProjectConfig()

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

func PrepareConfiguration() *CliArgsConfig {
	args, err := flags.Parse(CliConfig)
	if err != nil {
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				os.Exit(0)
			} else {
				log.Fatalf("Error parsing flags: %v", err)
				os.Exit(1)
			}
		default:
			log.Error(err)
			os.Exit(1)
		}
	}

	if len(args) == 0 && !CliConfig.StdinMode {
		log.Fatalf("No input file specified, usage should be: `arc ...options <relative file path>`")
	}

	if CliConfig.WorkingDirectory == "" {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatalf("Error getting working directory: %v", err)
		}
		CliConfig.WorkingDirectory = wd
	}

	if filepath.IsAbs(args[0]) {
		CliConfig.File = args[0]
	} else {
		CliConfig.File = filepath.Join(CliConfig.WorkingDirectory, args[0])
	}

	if CliConfig.File == "" && !CliConfig.StdinMode {
		log.Fatalf("No input file specified")
	}

	return CliConfig
}

func LoadProjectConfiguration() *ProjectConfiguration {
	configPath := filepath.Join(CliConfig.WorkingDirectory, "arc.json")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Warnf("No arc.json file found in working directory")
		return ProjectConfig
	}

	configContents, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Error reading arc.json file: %v", err)
	}

	var tmpConfig *ProjectConfiguration = &ProjectConfiguration{}
	if err := json.Unmarshal(configContents, tmpConfig); err != nil {
		log.Fatalf("Error marshaling project config: %v", err)
	}

	if tmpConfig.HttpServer == nil {
		tmpConfig.HttpServer = ProjectConfig.HttpServer
	} else {
		if tmpConfig.HttpServer.FormMaxMemory == 0 {
			tmpConfig.HttpServer.FormMaxMemory = ProjectConfig.HttpServer.FormMaxMemory
		}
		if tmpConfig.HttpServer.Port == nil {
			tmpConfig.HttpServer.Port = ProjectConfig.HttpServer.Port
		}
		if tmpConfig.HttpServer.Address == nil {
			tmpConfig.HttpServer.Address = ProjectConfig.HttpServer.Address
		}
		if tmpConfig.HttpServer.ReadHeaderTimeout == nil {
			tmpConfig.HttpServer.ReadHeaderTimeout = ProjectConfig.HttpServer.ReadHeaderTimeout
		}
		if tmpConfig.HttpServer.WriteTimeout == nil {
			tmpConfig.HttpServer.WriteTimeout = ProjectConfig.HttpServer.WriteTimeout
		}
	}

	ProjectConfig = tmpConfig

	if ProjectConfig.HttpServer.Port.EnvName != "" {
		port := os.Getenv(ProjectConfig.HttpServer.Port.EnvName)
		if port != "" {
			pv, err := strconv.Atoi(port)
			if err != nil {
				log.Fatalf("Error parsing port: %v", err)
			}
			ProjectConfig.HttpServer.Port.Value = pv
		} else {
			log.Warnf("No port specified in env variable: %v", ProjectConfig.HttpServer.Port.EnvName)
		}
	}

	if ProjectConfig.HttpServer.Address.EnvName != "" {
		address := os.Getenv(ProjectConfig.HttpServer.Address.EnvName)
		if address != "" {
			ProjectConfig.HttpServer.Address.Value = address
		} else {
			log.Warnf("No address specified in env variable: %v", ProjectConfig.HttpServer.Address.EnvName)
		}
	}

	return ProjectConfig
}
