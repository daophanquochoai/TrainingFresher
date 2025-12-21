package cache

import "github.com/google/wire"

var SetCache = wire.NewSet(NewRedisClient)
