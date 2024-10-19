package geometry

import "testing"

func TestPerimeter(t *testing.T) {
	cases := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{4.5, 0.8}, 10.6},
		{Circle{10}, 62.83185307179586},
	}

	for _, tt := range cases {
		checkPerimeter(t, tt.shape, tt.want)
	}
}

func checkPerimeter(t *testing.T, s Shape, want float64) {
	got := s.Perimeter()
	assert(t, got, want)
}
