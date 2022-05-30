package enums

import (
	"testing"
)

func TestEnumSet(t *testing.T) {
	set := NewEnumSet()
	set.RegOrDie("foo")
	set.RegOrDie("bar")
	// set.RegOrDie("foo")

	if !set.Verify("foo") {
		t.Errorf("TestEnumSet failed: Verify")
	}
}
