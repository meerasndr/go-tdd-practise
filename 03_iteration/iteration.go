package iteration

func Repeat(input string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += input
	}
	return repeated
}