package geometry

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("rectangle", func(t *testing.T) {
		r := Rectangle{4.5, 0.8}
		checkPerimeter(t, r, 10.6)
	})

	t.Run("circle", func(t *testing.T) {
		c := Circle{10}
		checkPerimeter(t, c, 62.83185307179586)
	})
}

func checkPerimeter(t *testing.T, s Shape, want float64) {
	got := s.Perimeter()
	assert(t, got, want)
}
