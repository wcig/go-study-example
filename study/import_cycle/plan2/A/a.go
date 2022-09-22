package A

import (
	"go-app/study/import_cycle/plan2/C"
	"strings"
)

type PackageA struct {
	B C.PackageBInterface
}

func (a *PackageA) Foo(s string) string {
	return a.B.Add(s)
}

func (a *PackageA) Trim(s string) string {
	return strings.Trim(s, "|")
}
