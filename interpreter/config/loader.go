package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strconv"

	"github.com/jessevdk/go-flags"

	"arc/log"
)

var CliConfig *CliArgsConfig = &CliArgsConfig{}
var ProjectConfig *ProjectConfiguration = createDefaultProjectConfig()

func PrepareConfiguration() *CliArgsConfig {
	flagParser := flags.NewParser(CliConfig, flags.Default)
	_, err := flagParser.Parse()
	if err != nil {

		var flagError *flags.Error
		if errors.As(err, &flagError) {
			// If we didn't print the help message, let's just print it
			if flagError.Type != flags.ErrHelp {
				flagParser.WriteHelp(os.Stdout)
				os.Exit(1)
			}
			os.Exit(0)
		}

		log.Fatalf("Error parsing command line arguments: %v", err)
	}

	if CliConfig.Extra.File == "" && !CliConfig.StdinMode {
		log.Fatalf("No input file specified, usage should be: `arc ...options <relative file path>`")
	}

	if CliConfig.WorkingDirectory == "" {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatalf("Error getting working directory: %v", err)
		}
		CliConfig.WorkingDirectory = wd
	}

	if CliConfig.Extra.File != "" {
		// CliConfig.WorkingDirectory
		if !filepath.IsAbs(CliConfig.Extra.File) {
			CliConfig.Extra.File = filepath.Join(CliConfig.WorkingDirectory, CliConfig.Extra.File)
		}
	}

	if CliConfig.Extra.File == "" && !CliConfig.StdinMode {
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
