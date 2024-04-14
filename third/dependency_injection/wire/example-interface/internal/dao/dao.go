package dao

import (
	"go-app/third/dependency_injection/wire/example-interface/internal/db"

	"github.com/google/wire"
)

var Provider = wire.NewSet(NewDao, wire.Bind(new(IDao), new(*Dao)))

type IDao interface {
	Version() (string, error)
}

type Dao struct {
	db *db.DB
}

func NewDao(db *db.DB) *Dao {
	return &Dao{db: db}
}

func (d *Dao) Version() (string, error) {
	var version string
	row := d.db.Mysql.QueryRow("SELECT VERSION()")
	err := row.Scan(&version)
	return version, err
}
