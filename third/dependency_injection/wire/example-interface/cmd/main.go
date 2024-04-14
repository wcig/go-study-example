package main

import (
	"go-app/third/dependency_injection/wire/example-interface/internal/dao"
	"log"
)

type App struct {
	dao dao.IDao
}

func NewApp(dao dao.IDao) (*App, error) {
	return &App{dao: dao}, nil
}

func main() {
	app, err := InitApp()
	if err != nil {
		log.Fatal(err)
	}

	version, err := app.dao.Version()
	log.Println(version, err)
}
