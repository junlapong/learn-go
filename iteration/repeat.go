package iteration

const repeatCount = 5

// Repeat character
func Repeat(character string) (repeated string) {
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return
}
