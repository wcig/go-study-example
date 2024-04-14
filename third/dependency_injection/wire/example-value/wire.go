//go:build wireinject
// +build wireinject

package main

import (
	"io"
	"os"

	"github.com/google/wire"
)

func InitApp() App {
	wire.Build(Provider, wire.Value(12), wire.InterfaceValue(new(io.Reader), os.Stdin), NewApp)
	return App{}
}
