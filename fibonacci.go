package tasks

func Fibonacci(limit int) []int {

	var fubonacciNums []int = []int{}

	var a, b, c int

	b = 1

	for i := 0; i < limit; i++ {

		fubonacciNums = append(fubonacciNums, c+b)
		a = c + b
		c = b
		b = a
	}

	return fubonacciNums
}
