package B

import (
	"fmt"
	"go-app/study/import_cycle/plan5/C"
)

func init() {
	C.RegisterFunc("add", Add)
}

func Bar(s string) string {
	trimFunc := C.GetFunc("trim")
	return trimFunc(s)
}

func Add(s string) string {
	return fmt.Sprintf("|%s|", s)
}
