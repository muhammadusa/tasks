package tasks

func FizzBuzz(number int) string {

	fizzbuzz := "fizzbuzz"
	fizz := "fizz"
	buzz := "buzz"

	if number%3 == 0 && number%5 == 0 {
		return fizzbuzz
	} else if number%3 == 0 {
		return fizz
	} else if number%5 == 0 {
		return buzz
	} else {
		return "Number isn't devide 3 and/ort 5"
	}
}
