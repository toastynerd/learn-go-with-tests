package array_and_slice

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) > 0 {
			sums = append(sums, Sum(numbers[1:]))
		} else {
			sums = append(sums, 0)
		}
	}
	return sums
}
