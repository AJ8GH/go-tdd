package geometry

import "testing"

func TestArea(t *testing.T) {
	cases := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{9.4, 7.3}, 68.62},
		{Rectangle{4, 2}, 8},
		{Circle{10}, 314.1592653589793},
	}

	for _, tt := range cases {
		checkArea(t, tt.shape, tt.want)
	}
}

func checkArea(t *testing.T, s Shape, want float64) {
	got := s.Area()
	assert(t, got, want)
}
