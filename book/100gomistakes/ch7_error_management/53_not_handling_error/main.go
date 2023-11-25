package main

import (
	"errors"
	"fmt"
)

func main() {
	// bad
	notify()

	// bad
	_ = notify()

	// good
	if err := notify(); err != nil {
		fmt.Println(err)
	}
}

func notify() error {
	return errors.New("failed to notify")
}
