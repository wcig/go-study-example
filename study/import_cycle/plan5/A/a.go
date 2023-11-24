package A

import (
	"go-app/study/import_cycle/plan5/C"
	"strings"
)

func init() {
	C.RegisterFunc("trim", Trim)
}

func Foo(s string) string {
	addFunc := C.GetFunc("add")
	return addFunc(s)
}

func Trim(s string) string {
	return strings.Trim(s, "|")
}
