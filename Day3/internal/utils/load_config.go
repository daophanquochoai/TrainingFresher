package utils

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// type config common
type Config struct {
	App AppConfig `yaml:"app"`
	DB  DBConfig  `yaml:"db"`
}

// type app conig
type AppConfig struct {
	Port int `yaml:"port"`
}

// type db config
type DBConfig struct {
	URL string `yaml:"url"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	log.Printf("Config loaded: %+v\n", config)
	return &config, nil
}
