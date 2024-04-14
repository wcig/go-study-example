package jsonx

import (
	"encoding/json"
	"fmt"
)

func PrintStr(v interface{}) {
	data, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(data))
}
