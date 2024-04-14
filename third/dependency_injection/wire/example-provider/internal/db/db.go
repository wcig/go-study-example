package db

import (
	"database/sql"
	"go-app/third/dependency_injection/wire/example-provider/internal/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
)

var Provider = wire.NewSet(NewDB)

type DB struct {
	Mysql *sql.DB
}

func NewDB(cfg *config.Config) (*DB, error) {
	db, err := sql.Open("mysql", cfg.Database.DSN)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	return &DB{Mysql: db}, nil
}
