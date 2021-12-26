package main

func SumArr(numbers [5]int) int {
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

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, SumSlice(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ... []int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		}else {
			tail := numbers[1:]
			sums = append(sums, SumSlice(tail))
		}
	}
	return sums
}