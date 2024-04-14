//go:build wireinject
// +build wireinject

package main

import (
	"go-app/third/dependency_injection/wire/example-project/internal/config"
	"go-app/third/dependency_injection/wire/example-project/internal/log"
	"go-app/third/dependency_injection/wire/example-project/internal/server"

	"github.com/google/wire"
)

func wireApp(*config.Config, *config.Logger, *config.Server) (*App, func(), error) {
	wire.Build(log.ProviderSet, server.ProviderSet, newApp)
	return &App{}, func() {}, nil
}
