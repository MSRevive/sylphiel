package dbot

import (
	"os"
	"errors"
	"path/filepath"

	"gopkg.in/ini.v1"
	"gopkg.in/yaml.v2"
	"github.com/disgoorg/snowflake/v2"
)

type Config struct {
	Core struct {
		Token string
		MaxThreads int
		DebugMode bool
	}
	Discord struct {
		DevChannel snowflake.ID
	}
	Log struct {
		Level string
		Dir string
		ExpireTime string
	}
}

func LoadConfig(path string, dbg bool) (*Config, error) {
	var cfg Config

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