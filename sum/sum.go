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

func SumAll(ints [][]int) []int {
	var sums []int
	for _, s := range ints {
		sums = append(sums, SumSlice(s))
	}
	return sums
}

func SumAllTails(ints [][]int) []int {
	var out []int
	for _, s := range ints {
		out = append(out, s[len(s)-1])
	}
	return out
}
