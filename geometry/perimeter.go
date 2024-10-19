package geometry

func Perimeter(r Rectangle) float64 {
	return (r.Height + r.Width) * 2
}

func Area(r Rectangle) float64 {
	return r.Height * r.Width
}
