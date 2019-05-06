package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if want != got {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2}, []int{0, 9})
	want := []int{2, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSlice(t *testing.T) {
	numbers := []int{0, 1, 2, 3, 4}
	fmt.Printf("%v => len: %d, cap: %d\n", numbers, len(numbers), cap(numbers))
	fmt.Printf("%v\n", numbers[1:])
	fmt.Printf("%v\n", numbers[1:3])

	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	fmt.Printf("%c\n", b)
	fmt.Printf("%c\n", b[:2])
	fmt.Printf("%c\n", b[1:4])

	// TODO: len, cap
}
