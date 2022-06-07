package arrays

import (
	"reflect"
	"testing"
)

var inputArray = func() [5]int {
	return [5]int{1, 2, 3, 4, 5}
}

var inputSlice = func() []int {
	return []int{1, 2, 3}
}

var inputSlices = func() ([]int, []int) {
	return []int{1, 2, 3}, []int{4, 5, 6}
}

var inputEmptySlices = func() ([]int, []int) {
	return []int{}, []int{4, 5, 6}
}

var assertArray = func(t testing.TB, got, want int, input [5]int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, input)
	}
}

var assertSlice = func(t testing.TB, got, want int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

var assertSliceAll = func(t testing.TB, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSum(t *testing.T) {
	t.Run("sum an array using classic for", func(t *testing.T) {
		input := inputArray()
		got := SumClassic(input)
		want := 15

		assertArray(t, got, want, input)
	})
	t.Run("sum an array using range", func(t *testing.T) {
		input := inputArray()
		got := SumRange(input)
		want := 15

		assertArray(t, got, want, input)
	})
	t.Run("sum a slice", func(t *testing.T) {
		input := inputSlice()
		got := SumSlice(input)
		want := 6

		assertSlice(t, got, want)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("sum and return slices classic", func(t *testing.T) {
		got := SumAllClassic(inputSlices())
		want := []int{6, 15}
		assertSliceAll(t, got, want)
	})
	t.Run("sum and return slices using append", func(t *testing.T) {
		got := SumAllAppend(inputSlices())
		want := []int{6, 15}
		assertSliceAll(t, got, want)
	})
}

func TestSumAllTail(t *testing.T) {
	t.Run("sum tail of slices and return as slice", func(t *testing.T) {
		got := SumAllTail(inputSlices())
		want := []int{5, 11}

		assertSliceAll(t, got, want)
	})
	t.Run("sum tail of empty slices", func(t *testing.T) {
		got := SumAllTail(inputEmptySlices())
		want := []int{0, 11}

		assertSliceAll(t, got, want)
	})
}
