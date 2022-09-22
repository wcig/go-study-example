package C

import (
	"go-app/study/import_cycle/plan3/A"
	"go-app/study/import_cycle/plan3/B"
)

type PackageC struct {
	A *A.PackageA
	B *B.PackageB
}

func (c *PackageC) FooAdd(s string) string {
	return c.B.Add(s)
}

func (c *PackageC) BarTrim(s string) string {
	return c.A.Trim(s)
}
