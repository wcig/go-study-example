//go:build wireinject
// +build wireinject

package main

import (
	"go-app/third/dependency_injection/wire/example-provider/internal/config"
	"go-app/third/dependency_injection/wire/example-provider/internal/db"

	"github.com/google/wire"
)

func InitApp() (*App, error) {
	wire.Build(config.Provider, db.Provider, NewApp)
	return &App{}, nil
}
