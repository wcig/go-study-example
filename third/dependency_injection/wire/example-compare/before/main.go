package main

import "log"

func main() {
	cfg := NewConfig()
	db := NewDB(cfg)
	result := db.Ping()
	log.Println(result)
}
