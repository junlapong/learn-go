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

// https://gobyexample.com/slices
func TestSlice(t *testing.T) {

	s := make([]string, 3)
	fmt.Println("emp:", s)

	// We can set and get just like with arrays.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// `len` returns the length of the slice as expected.
	fmt.Println("len:", len(s))

	// In addition to these basic operations, slices
	// support several more that make them richer than
	// arrays. One is the builtin `append`, which
	// returns a slice containing one or more new values.
	// Note that we need to accept a return value from
	// `append` as we may get a new slice value.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Slices can also be `copy`'d. Here we create an
	// empty slice `c` of the same length as `s` and copy
	// into `c` from `s`.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Slices support a "slice" operator with the syntax
	// `slice[low:high]`. For example, this gets a slice
	// of the elements `s[2]`, `s[3]`, and `s[4]`.
	l := s[2:5]
	fmt.Printf("s      : %v, len: %d, cap: %d\n", s, len(s), cap(s))
	fmt.Println("s[2:5] :", l)

	// This slices up to (but excluding) `s[5]`.
	l = s[:5]
	fmt.Println("s[:5]  :", l)

	// And this slices up from (and including) `s[2]`.
	l = s[2:]
	fmt.Println("s[2:]  :", l)

	// We can declare and initialize a variable for slice
	// in a single line as well.
	st := []string{"g", "h", "i"}
	fmt.Println("dcl:", st)

	// Slices can be composed into multi-dimensional data
	// structures. The length of the inner slices can
	// vary, unlike with multi-dimensional arrays.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}

func TestSliceOperator(t *testing.T) {

	a := [5]int{1, 2, 3, 4, 5}
	b := a[0:3]

	fmt.Println(a)
	fmt.Println(b)
	// b จะกลายเป็นขอมูลแบบ slice เพราะเกิดจากการ slice array a
	// จากตำแหน่งที่ 0 จนถึงตำแหน่งก่อนหน้า 3 ดังนั้นข้อมูลของ b คือ [1,2,3]

	fmt.Println(a[1:5])
	fmt.Println(a[2:4])
}

func TestArraySlice(t *testing.T) {

	var a []int
	a = append(a, 1)
	fmt.Printf("a: %v, len: %d, cap: %d\n", a, len(a), cap(a))

	a = []int{1, 2, 3, 4}
	fmt.Printf("a: %v\n", a)

	b := a

	c := make([]int, len(a))
	copy(c, a)

	d := append([]int(nil), a...)

	// change b[0] value
	b[0] = 5

	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
	fmt.Printf("d: %v\n", d)

}
