package array_and_slice

func SumAllHeads(numbersToSum... []int) int {
	sum := 0
	for _, numbers := range numbersToSum {
		if len(numbers) > 0 {
			sum += numbers[0]
		}
	}
	return sum
}
