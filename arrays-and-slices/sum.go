package sum

func Sum(numbers []int) int {
	sum := 0

	// classic iteration
	//for i := 0; i < len(numbers); i++ {
	//	sum += numbers[i]
	//}

	// using range
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	//listsCount := len(numbersToSum)
	//sum := make([]int, listsCount)
	//
	//for index, row := range numbersToSum {
	//	sum[index] = Sum(row)
	//}

	// refactoring
	var sum []int

	for _, row := range numbersToSum {
		sum = append(sum, Sum(row)) // => sum.push(Sum(row))
	}

	return sum
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sum []int

	for _, row := range numbersToSum {
		// "take from 1 to the end"
		tail := row[1:] // slices, pretty the same as in python (i guess). syntax: [low:high]. [1:] will return all items from index 1 to last one in the slice/array
		sum = append(sum, Sum(tail))
	}

	return sum
}
