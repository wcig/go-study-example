package xorm

import (
	"testing"
)

var (
	host   = "10.200.50.18"
	port   = "3306"
	user   = "root"
	passwd = "123456"
	db     = "test"
)

func TestAutoGen(t *testing.T) {
	dir := "out"
	AutoGen(user, passwd, host, port, db, dir)
}
