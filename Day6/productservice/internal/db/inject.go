package db

import "github.com/google/wire"

var Set = wire.NewSet(Connect)
