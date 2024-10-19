package sum

import (
	"reflect"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("array", func(t *testing.T) {
		ints := [...]int{1, 2, 3, 4, 5}
		got := Sum(ints)
		want := 15
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("slice", func(t *testing.T) {
		ints := []int{1, 2, 3, 4, 5}
		got := SumSlice(ints)
		want := 15
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("many slices", func(t *testing.T) {
		ints := [][]int{{1, 2, 3}, {5, 5}, {7, 9}}
		got := SumAll(ints)
		want := []int{6, 10, 16}
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("all tails", func(t *testing.T) {
		ints := [][]int{{1, 2, 3}, {5, 5}, {7, 9}}
		got := SumAllTails(ints)
		want := []int{3, 5, 9}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
