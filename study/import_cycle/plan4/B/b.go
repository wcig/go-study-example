package B

import (
	"fmt"
	"go-app/study/import_cycle/plan4/C"
)

func Bar(s string) string {
	return C.TrimFunc(s)
}

func Add(s string) string {
	return fmt.Sprintf("|%s|", s)
}
