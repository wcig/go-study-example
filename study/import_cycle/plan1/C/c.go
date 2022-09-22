package C

import (
	"fmt"
	"strings"
)

func Trim(s string) string {
	return strings.Trim(s, "|")
}

func Add(s string) string {
	return fmt.Sprintf("|%s|", s)
}
