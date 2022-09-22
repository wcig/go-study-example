package B

import (
	"fmt"
	"go-app/study/import_cycle/plan2/C"
)

type PackageB struct {
	A C.PackageAInterface
}

func (b *PackageB) Bar(s string) string {
	return b.A.Trim(s)
}

func (b *PackageB) Add(s string) string {
	return fmt.Sprintf("|%s|", s)
}
