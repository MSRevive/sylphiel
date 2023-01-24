package system

import (
	"os"
	"path/filepath"
	"errors"

	"gopkg.in/ini.v1"
	"gopkg.in/yaml.v2"
)

var (
	Version = "canary"
)

type config struct {
	Core struct {
		Token string
		MaxThreads int
		DebugMode bool
	}
	Log struct {
		Level string
		Dir string
		ExpireTime string
	}
}

func LoadConfig(path string, dbg bool) (*config, error) {
	var cfg config

	switch filepath.Ext(path) {
	case ".toml", ".ini":
		if err := ini.MapTo(&cfg, path); err != nil {
			return nil, err
		}

	case ".yaml", ".json", ".yml":
		data,err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		
		if err := yaml.Unmarshal(data, &cfg); err != nil {
			return nil, err
		}

	default:
		return nil, errors.New("unsupported config type")
	}

	cfg.Core.DebugMode = dbg
	return &cfg, nil
}