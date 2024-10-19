package geometry

import "testing"

func TestArea(t *testing.T) {
	t.Run("rectangle", func(t *testing.T) {
		r := Rectangle{9.4, 7.3}
		checkArea(t, r, 68.62)
	})

	t.Run("circle", func(t *testing.T) {
		c := Circle{10}
		checkArea(t, c, 314.1592653589793)
	})
}

func checkArea(t *testing.T, s Shape, want float64) {
	got := s.Area()
	assert(t, got, want)
}
