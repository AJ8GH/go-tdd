package geometry

import "testing"

func TestPerimeter(t *testing.T) {
	r := Rectangle{4.5, 0.8}
	got := Perimeter(r)
	want := 10.6
	assert(t, got, want)
}

func TestArea(t *testing.T) {
	r := Rectangle{9.4, 7.3}
	got := Area(r)
	want := 68.62
	assert(t, got, want)
}

func TestAreaCircle(t *testing.T) {
	c := Circle{10}
	got := c.Area()
	want := 314.1592653589793
	assert(t, got, want)
}
func assert(t *testing.T, got, want float64) {
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}
