package arrays

// Sum nummers
func Sum(numbers []int) (sum int) {

	for _, number := range numbers {
		sum += number
	}

	return
}
