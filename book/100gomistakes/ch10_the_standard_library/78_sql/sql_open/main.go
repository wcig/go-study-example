package main

import "database/sql"

var dsn = ""

func listing1() error {
	// 只创建一个实例对象, 没有建立tcp连接
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	// 此时才真正建立tcp连接
	if err = db.Ping(); err != nil {
		return err
	}

	_ = db
	return nil
}
