package sum

import (
	"reflect"
	"testing"
)

func checkSums[T any](t testing.TB, got, want T) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSum(t *testing.T) {
	t.Run("array", func(t *testing.T) {
		ints := [...]int{1, 2, 3, 4, 5}
		got := Sum(ints)
		want := 15
		checkSums(t, got, want)
	})

	t.Run("slice", func(t *testing.T) {
		ints := []int{1, 2, 3, 4, 5}
		got := SumSlice(ints)
		want := 15
		checkSums(t, got, want)
	})

	t.Run("many slices", func(t *testing.T) {
		ints := [][]int{{1, 2, 3}, {5, 5}, {7, 9}}
		got := SumAll(ints)
		want := []int{6, 10, 16}
		checkSums(t, got, want)
	})

	t.Run("all tails", func(t *testing.T) {
		ints := [][]int{{1, 2, 3}, {5, 5}, {7, 9}}
		got := SumAllTails(ints)
		want := []int{3, 5, 9}
		checkSums(t, got, want)
	})

	t.Run("all tails with empties", func(t *testing.T) {
		ints := [][]int{{}, {1, 2, 1}, {}, {5, 2}, {7, 6}}
		got := SumAllTails(ints)
		want := []int{1, 2, 6}
		checkSums(t, got, want)
	})
}
