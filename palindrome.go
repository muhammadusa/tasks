package tasks

func Palindrome(number int) string {

	summa, tmp := 0, number

	for tmp > 0 {

		rem := tmp % 10

		summa = summa*10 + rem

		tmp /= 10

	}

	if summa == number {
		return "Palindrome"
	} else {
		return "Not Palindrome"
	}
}
