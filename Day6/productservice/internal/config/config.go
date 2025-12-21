package config

import (
	"fmt"
	"strings"
	"github.com/spf13/viper"
)

type Config struct {
	App   App   `mapstructure:"app"`
	DB    DB    `mapstructure:"db"`
	Redis Redis `mapstructure:"redis"`
}

type App struct {
	Name string `mapstructure:"name"`
	Grpc string `mapstructure:"grpc"`
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
	DB       int    `mapstructure:"db"` // ✅ Sửa từ db string -> DB int
}

func LoadConfig(path string) (*Config, error) {
	v := viper.New()

	// config file
	v.SetConfigFile(path)
	v.SetConfigType("yaml")

	// ENV support - bind environment variables for nested structures
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	
	// Bind environment variables explicitly
	v.BindEnv("app.name", "APP_NAME")
	v.BindEnv("app.grpc", "APP_GRPC")
	v.BindEnv("app.port", "APP_PORT")
	v.BindEnv("db.host", "DB_HOST")
	v.BindEnv("db.port", "DB_PORT")
	v.BindEnv("db.user", "DB_USER")
	v.BindEnv("db.pass", "DB_PASS")
	v.BindEnv("db.name", "DB_NAME")
	v.BindEnv("redis.host", "REDIS_HOST")
	v.BindEnv("redis.pass", "REDIS_PASSWORD")
	v.BindEnv("redis.db", "REDIS_DB")

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

func NewConfig() (*Config, error) {
	configPath := "./config/config.yaml"
	return LoadConfig(configPath)
}

func NewDBConfig(cfg *Config) DB {
	return cfg.DB
}

func NewRedisConfig(cfg *Config) Redis {
	return cfg.Redis
}

func NewAppConfig(cfg *Config) App {
	return cfg.App
}
