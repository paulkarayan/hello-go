package main

func Sum(numbers []int) int {
	sum := 0

	// for i := 0; i < 5; i++ {
	// 	sum += numbers[i]
	// }

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	// lengthOfNumbers := len(numbersToSum)
	// // make a slice of ints, the length of the numbers we have to sum
	// sums := make([]int, lengthOfNumbers)

	// for i, numbers := range numbersToSum {
	// 	sums[i] = Sum(numbers)
	// }

	// version not worrying about size

	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

// sum of all but the first item (the "head")
func SumAllTails(numbersToSum ...[]int) []int {

	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}

	}

	return sums
}
