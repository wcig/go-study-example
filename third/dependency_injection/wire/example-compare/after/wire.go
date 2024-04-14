//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func InitApp() (*App, error) {
	wire.Build(NewConfig, NewDB, NewApp)
	return &App{}, nil
}
