package sum

func Sum(ints [5]int) int {
	sum := 0
	for _, n := range ints {
		sum += n
	}
	return sum
}

func SumSlice(ints []int) int {
	sum := 0
	for _, n := range ints {
		sum += n
	}
	return sum
}

func SumAll(ints [][]int) int {
	sum := 0
	for _, n := range ints {
		sum += SumSlice(n)
	}
	return sum
}
