package tasks

func OddEvenSum(numbers []int) (oddSum, evenSum int) {

	for _, number := range numbers {

		if number%2 == 0 {
			evenSum += number
		} else {
			oddSum += number
		}
	}
	return oddSum, evenSum
}
