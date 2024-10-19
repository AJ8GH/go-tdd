package geometry

import "testing"

func TestArea(t *testing.T) {
	cases := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{
			name:  "Rectangle-1",
			shape: Rectangle{Height: 9.4, Width: 7.3},
			want:  68.62,
		},
		{
			name:  "Rectangle-2",
			shape: Rectangle{Height: 4, Width: 2},
			want:  8,
		},
		{
			name:  "Circle",
			shape: Circle{Radius: 10},
			want:  314.1592653589793,
		},
		{
			name:  "Triangle",
			shape: Triangle{A: 2, B: 2, C: 2},
			want:  1.7320508075688772,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			checkArea(t, tt.shape, tt.want)
		})
	}
}

func checkArea(t *testing.T, s Shape, want float64) {
	got := s.Area()
	assert(t, s, got, want)
}
