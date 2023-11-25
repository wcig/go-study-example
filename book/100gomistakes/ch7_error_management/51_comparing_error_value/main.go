package main

import (
	"database/sql"
	"errors"
	"fmt"
)

func main() {
	err := fmt.Errorf("warp err: %w", sql.ErrNoRows)

	// bad
	fmt.Println(err == sql.ErrNoRows) // false

	// good
	fmt.Println(errors.Is(err, sql.ErrNoRows)) // true
}
