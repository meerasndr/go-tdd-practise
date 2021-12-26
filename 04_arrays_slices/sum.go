package main
func SumArr(numbers [5] int) int{
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumSlice(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}