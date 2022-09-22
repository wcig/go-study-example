package B

import (
	"fmt"
	"go-app/study/import_cycle/question/A"
)

func Bar(s string) string {
	return A.Trim(s)
}

func Add(s string) string {
	return fmt.Sprintf("|%s|", s)
}
