package main

import (
	"go-app/third/dependency_injection/wire/example-provider/internal/db"
	"log"
)

type App struct {
	db *db.DB
}

func NewApp(db *db.DB) (*App, error) {
	return &App{db: db}, nil
}

func main() {
	app, err := InitApp()
	if err != nil {
		log.Fatal(err)
	}

	var version string
	row := app.db.Mysql.QueryRow("SELECT VERSION()")
	if err = row.Scan(&version); err != nil {
		log.Fatal(err)
	}
	log.Println("mysql version:", version)
}
