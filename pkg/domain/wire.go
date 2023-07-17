//go:build wireinject
// +build wireinject

package domain

import (
	"github.com/citrusinesis/thread/pkg/config"
	"github.com/citrusinesis/thread/pkg/domain/entity"
	"github.com/google/wire"
)

func InjectRouter() *Router {
	panic(wire.Build(
		config.Set,
		entity.Set,
		NewRouter,
	))
}
