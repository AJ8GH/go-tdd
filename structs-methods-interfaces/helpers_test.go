package geometry

import "testing"

func assert(t testing.TB, shape Shape, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("%#v got %g want %g", shape, got, want)
	}
}
