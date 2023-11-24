package main

import (
	"fmt"
	"go-app/book/100gomistakes/ch2_code_project_organization/3_init_functions/good/db"
	"go-app/book/100gomistakes/ch2_code_project_organization/3_init_functions/good/redis"
)

func init() {
	db.InitDB()
	redis.InitRedis()
	initMain()
}

func initMain() {
	fmt.Println(">> init main")
}

func main() {
	// Output:
	// >> init db
	// >> init redis
	// >> init main
}
