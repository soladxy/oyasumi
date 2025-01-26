package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

func NewServiceConfig(path string) (*Config, error) {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	config := &Config{}
	err = yaml.Unmarshal(fileBytes, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

type Config struct {
	MySQL MySQLConfig `yaml:"MySQL"`
}

type MySQLConfig struct {
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	Host     string `yaml:"Host"`
	Port     int    `yaml:"Port"`
}
