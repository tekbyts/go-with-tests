package arrays

func SumClassic(numbers [5]int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return sum
}

func SumRange(numbers [5]int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumSlice(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAllClassic(numbers ...[]int) (sums []int) {
	sums = make([]int, len(numbers))
	for i, set := range numbers {
		sums[i] = SumSlice(set)
	}
	return
}

func SumAllAppend(numbers ...[]int) []int {
	var sum []int
	for _, set := range numbers {
		sum = append(sum, SumSlice(set))
	}
	return sum
}

func SumAllTail(numbers ...[]int) (sum []int) {
	for _, set := range numbers {
		if len(set) > 0 {
			sum = append(sum, SumSlice(set[1:]))
		} else {
			sum = append(sum, 0)
		}
	}
	return
}
