package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	err := listing1()
	fmt.Println(err) // unexpected end of JSON input
}

func listing1() error {
	b := getMessage()
	var m map[string]any
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	return nil
}

func getMessage() []byte {
	return nil
}
