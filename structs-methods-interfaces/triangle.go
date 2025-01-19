package geometry

import "math"

type Triangle struct {
	A float64
	B float64
	C float64
}

func (t Triangle) Area() float64 {
	return 0.25 * math.Sqrt(
		4*math.Pow(t.A, 2)*math.Pow(t.B, 2)-
			math.Pow(math.Pow(t.A, 2)+math.Pow(t.B, 2)-math.Pow(t.C, 2), 2))
}

func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}
