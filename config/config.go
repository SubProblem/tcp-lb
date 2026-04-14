package config

import (
	"os"
	"gopkg.in/yaml.v3"
)

type Config struct {
	ListenAddr string   `yaml:"listen_addr"`
	Backends   []string `yaml:"backends"`
	Strategy   string   `yaml:"strategy"`
	StatsAddr  string   `yaml:"stats_addr"`
}

func (c *Config) LoadConfig() error {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, c)
}