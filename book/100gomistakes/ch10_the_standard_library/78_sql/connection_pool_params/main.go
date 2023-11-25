package main

import (
	"database/sql"
	"time"
)

var dsn = ""

func listing1() error {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(30)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(time.Hour)
	db.SetConnMaxLifetime(time.Hour)

	if err = db.Ping(); err != nil {
		return err
	}

	_ = db
	return nil
}
