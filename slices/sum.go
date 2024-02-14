package slices

func Sum(array []int) int {
	sum := 0
	for _, number := range array {
		sum += number
	}
	return sum
}

func SumAll(slices ...[]int) []int {
	var sums []int

	for _, numbers := range slices {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(slices ...[]int) []int {
	var sums []int

	for _, numbers := range slices {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(numbers[1:]))
		}
	}

	return sums
}