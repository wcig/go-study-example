// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

// Injectors from wire.go:

func InitApp() (*App, error) {
	config := NewConfig()
	db := NewDB(config)
	app := NewApp(db)
	return app, nil
}
