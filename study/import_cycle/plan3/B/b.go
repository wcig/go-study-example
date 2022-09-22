package B

import (
	"fmt"
)

type PackageB struct{}

func (b *PackageB) Add(s string) string {
	return fmt.Sprintf("|%s|", s)
}
