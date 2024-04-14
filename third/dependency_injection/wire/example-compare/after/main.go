package main

import "log"

type App struct {
	db *DB
}

func NewApp(db *DB) *App {
	return &App{
		db: db,
	}
}

func main() {
	app, err := InitApp()
	if err != nil {
		log.Fatal(err)
	}
	result := app.db.Ping()
	log.Println(result)
}
