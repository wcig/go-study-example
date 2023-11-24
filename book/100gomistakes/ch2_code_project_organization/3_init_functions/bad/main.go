package main

import (
	"fmt"
	_ "go-app/book/100gomistakes/ch2_code_project_organization/3_init_functions/bad/db"
	_ "go-app/book/100gomistakes/ch2_code_project_organization/3_init_functions/bad/redis"
)

func main() {
	fmt.Println(">> init main")

	// Output:
	// >> init db
	// >> init redis
	// >> init main
}
