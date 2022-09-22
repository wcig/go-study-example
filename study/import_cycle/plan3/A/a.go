package A

import (
	"strings"
)

type PackageA struct{}

func (a *PackageA) Trim(s string) string {
	return strings.Trim(s, "|")
}
