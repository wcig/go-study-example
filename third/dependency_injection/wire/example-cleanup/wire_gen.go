// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

// Injectors from wire.go:

func InitApp() (*App, func(), error) {
	logger, cleanup, err := NewLogger()
	if err != nil {
		return nil, nil, err
	}
	app := NewApp(logger)
	return app, func() {
		cleanup()
	}, nil
}
