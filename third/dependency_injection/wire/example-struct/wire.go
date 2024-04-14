//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func NewFooBar() *FooBar {
	wire.Build(ProviderFoo, ProviderBar, wire.Struct(new(FooBar), "MyFoo", "MyBar"))
	// wire.Build(ProviderFoo, ProviderBar, wire.Struct(new(FooBar), "*")) // 效果一样
	return &FooBar{}
}
