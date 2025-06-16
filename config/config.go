package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

type RodConfig struct {
	Headless bool `yaml:"headless"`
	Timeout  int  `yaml:"timeout"`
}

type BrowserConfig struct {
	BrowserPath string `yaml:"browser_path"`
	Viewport    string `yaml:"viewport"`
}

type NetworkConfig struct {
	Throttle string `yaml:"throttle"`
}

type Config struct {
	Rod     RodConfig     `yaml:"rod"`
	Browser BrowserConfig `yaml:"browser"`
	Network NetworkConfig `yaml:"network"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	return &cfg, err
}

func (c *RodConfig) TimeoutDuration() time.Duration {
	return time.Duration(c.Timeout) * time.Second
}
