package A

import (
	"go-app/study/import_cycle/question/B"
	"strings"
)

func Foo(s string) string {
	return B.Add(s)
}

func Trim(s string) string {
	return strings.Trim(s, "|")
}
