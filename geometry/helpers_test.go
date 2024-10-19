package geometry

import "testing"

func assert(t *testing.T, shape Shape, got, want float64) {
	if got != want {
		t.Errorf("%#v got %g want %g", shape, got, want)
	}
}
