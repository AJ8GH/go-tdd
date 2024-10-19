package geometry

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(4.5, 0.8)
	want := 10.6
	assert(t, got, want)
}

func assert(t *testing.T, got, want float64) {
	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}
