package tasks

func HasDublicate(numbers []int) string {

	var dublicate bool

	for i, number := range numbers {

		for j := i; j < len(numbers)-1; j++ {

			if number == numbers[j+1] {
				dublicate = true
				break
			}
		}
	}
	if dublicate {

		return "Array has duplicate element"
	} else {

		return "Array has not duplicate element"
	}

}
