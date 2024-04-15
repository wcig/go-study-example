package jsonx

import (
	"encoding/json"
	"fmt"
)

func PrintStr(v interface{}) {
	fmt.Println(ToStr(v))
}

func ToStr(v interface{}) string {
	data, _ := json.MarshalIndent(v, "", "  ")
	return string(data)
}
