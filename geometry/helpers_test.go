package geometry

import "testing"

func assert(t *testing.T, got, want float64) {
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}
