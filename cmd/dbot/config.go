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
	}
	Disc struct {
		GuildID snowflake.ID
		DevChannel snowflake.ID
		RoleChannel snowflake.ID
	}
	Webhook struct {
		Enabled bool
		ID snowflake.ID
		Token string
	}
	Log struct {
		Level string
		Dir string
		ExpireTime string
	}
	FuzzNet struct {
		Domain string
		APIKey string
	}
}

func LoadConfig(path string, dbg bool) (*Config, error) {
	var cfg Config

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return nil, os.ErrNotExist
	}

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

	return &cfg, nil
}