package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	app, cleanup, err := InitApp()
	if err != nil {
		panic(err)
	}
	defer cleanup()
	app.logger.Println("app run")
}

type App struct {
	logger *log.Logger
}

func NewApp(logger *log.Logger) *App {
	return &App{logger: logger}
}

func NewLogger() (*log.Logger, func(), error) {
	file := os.Stdout
	logger := log.New(file, "App: ", log.Ldate|log.Ltime|log.Lshortfile)
	cleanup := func() {
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}
	return logger, cleanup, nil
}
