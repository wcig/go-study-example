package A

import (
	"go-app/study/import_cycle/plan4/C"
	"strings"
)

func Foo(s string) string {
	return C.AddFunc(s)
}

func Trim(s string) string {
	return strings.Trim(s, "|")
}
