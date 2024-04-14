//go:build wireinject
// +build wireinject

package main

import (
	"go-app/third/dependency_injection/wire/example-interface/internal/config"
	"go-app/third/dependency_injection/wire/example-interface/internal/dao"
	"go-app/third/dependency_injection/wire/example-interface/internal/db"

	"github.com/google/wire"
)

func InitApp() (*App, error) {
	wire.Build(config.Provider, db.Provider, dao.Provider, NewApp)
	return &App{}, nil
}
