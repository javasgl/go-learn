package simplemath

import (
	"testing"
)

func TestSqrt2(t *testing.T) {
	r := Sqrt(16)
	if r != 4 {
		t.Errorf("Sqrt(9) failed,Got %v,expected 3.", r)
	}
}
