package config

import (
	"os"

	"github.com/gosidekick/goconfig"
	_ "github.com/gosidekick/goconfig/json"
)

type Config struct {
	Output           string   `json:"output" cfg:"output" cfgDefault:"-"`
	ListenAddress    string   `json:"listen_address" cfg:"listen_address" cfgDefault:"0.0.0.0:8765"`
	Interval         int      `json:"interval" cfg:"interval" cfgDefault:"5"`
	OutputDescriptor *os.File `json:"-" cfg:"-"`
}

func Load() (*Config, error) {

	cfg := &Config{}

	goconfig.File = "config.json"
	err := goconfig.Parse(cfg)
	if err != nil {
		return nil, err
	}

	cfg.OutputDescriptor = os.Stdout
	return cfg, nil
}
