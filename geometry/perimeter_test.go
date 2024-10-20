package geometry

import "testing"

func TestPerimeter(t *testing.T) {
	cases := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{
			name:  "Rectangle",
			shape: Rectangle{Height: 4.5, Width: 0.8},
			want:  10.6},
		{
			name:  "Circle",
			shape: Circle{Radius: 10},
			want:  62.83185307179586},
		{
			name:  "Triangle",
			shape: Triangle{A: 2, B: 2, C: 2},
			want:  6,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			checkPerimeter(t, tt.shape, tt.want)
		})
	}
}

func checkPerimeter(t testing.TB, s Shape, want float64) {
	t.Helper()
	got := s.Perimeter()
	assert(t, s, got, want)
}
