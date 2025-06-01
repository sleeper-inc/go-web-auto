package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

type BrowserConfig struct {
	Headless    bool   `yaml:"headless"`
	Timeout     int    `yaml:"timeout"`
	BrowserPath string `yaml:"browser_path"`
}

type Config struct {
	Browser BrowserConfig `yaml:"browser"`
}

func LoadConfig(path string) (*Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(b, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (c *BrowserConfig) TimeoutDuration() time.Duration {
	return time.Duration(c.Timeout) * time.Second
}
