package counter

import (
	"testing"
)

func TestCount(t *testing.T) {
	if Inc() != 1 {
		t.Errorf("expected 1")
	}
}
