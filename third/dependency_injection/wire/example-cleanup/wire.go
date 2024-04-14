//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func InitApp() (*App, func(), error) {
	wire.Build(NewLogger, NewApp)
	return &App{}, func() {}, nil
}
