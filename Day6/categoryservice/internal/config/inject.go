package config

import "github.com/google/wire"

var Set = wire.NewSet(
	NewConfig,
	NewDBConfig,
	NewRedisConfig,
	NewConfigApp,
)
