package arrays

import "fmt"

// Sum nummers
func Sum(numbers []int) (sum int) {

	for _, number := range numbers {
		sum += number
	}

	return
}

// SumAll numbers
func SumAll(numbersToSum ...[]int) (sums []int) {

	// sums = make([]int, len(numbersToSum))
	// for i, numbers := range numbersToSum {
	// 	sums[i] = Sum(numbers)
	// }

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return
}

// SumAllTails numbers
func SumAllTails(numbersToSum ...[]int) (sums []int) {

	for _, numbers := range numbersToSum {
		tail := numbers[1:]
		fmt.Printf("tail: %v\n", tail)
		sums = append(sums, Sum(tail))
	}

	return
}
