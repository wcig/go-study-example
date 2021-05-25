package xorm

import (
	"testing"
)

var (
	host     = "localhost"
	port     = "3306"
	username = "root"
	password = "123456"
	dbname   = "test"
	dir      = "out"
)

func TestAutoGen(t *testing.T) {
	AutoGen(username, password, host, port, dbname, dir)
}
