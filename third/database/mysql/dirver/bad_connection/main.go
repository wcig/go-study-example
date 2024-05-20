package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// 配置mysql server的wait_timeout为10秒
var (
	db  *sql.DB
	err error
)

func main() {
	initDB()
	time.Sleep(time.Second * 12)
	query()
	// Output:
	// [mysql] 2022/08/08 18:18:37 packets.go:123: closing bad idle connection: EOF
	// query: driver: bad connection
}

func query() {
	if err := db.Ping(); err != nil {
		fmt.Println("query:", err)
	}
}

func initDB() {
	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute)
	if err = db.Ping(); err != nil {
		panic(err)
	}
}
