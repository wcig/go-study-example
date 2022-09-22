package B

import "go-app/study/import_cycle/plan1/C"

func Bar(s string) string {
	return C.Trim(s)
}
