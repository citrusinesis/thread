package config

import "github.com/google/wire"

// Set Dependency injection target in config
var Set = wire.NewSet(GetMySQLClient, NewDBConfig)
