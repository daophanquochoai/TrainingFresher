package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	App   App   `mapstructure:"app"`
	DB    DB    `mapstructure:"db"`
	Redis Redis `mapstructure:"redis"`
}

type App struct {
	Name string `mapstructure:"name"`
	Grpc bool   `mapstructure:"grpc"`
	Port int    `mapstructure:"port"`
}

type DB struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	User string `mapstructure:"user"`
	Pass string `mapstructure:"pass"`
	Name string `mapstructure:"name"`
}

type Redis struct {
	Addr     string `mapstructure:"host"`
	Password string `mapstructure:"pass"`
	db       string `mapstructure:"db"`
}

func LoadConfig(path string) (*Config, error) {
	v := viper.New()

	// config file
	v.SetConfigFile(path)
	v.SetConfigType("yaml")

	// ENV support
	v.AutomaticEnv()

	// doc file
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config error: %w", err)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config error: %w", err)
	}

	return &cfg, nil
}
