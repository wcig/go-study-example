package A

import "go-app/study/import_cycle/plan1/C"

func Foo(s string) string {
	return C.Add(s)
}
