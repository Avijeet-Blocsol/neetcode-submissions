func twoSum(numbers []int, target int) []int {
	rightIndex := len(numbers) - 1
	leftIndex := 0

	for leftIndex < rightIndex {
		if numbers[leftIndex]+numbers[rightIndex] == target {
			return []int{leftIndex + 1, rightIndex + 1}
		}

		if numbers[leftIndex]+numbers[rightIndex] < target {
			leftIndex++
		} else {
			rightIndex--
		}
	}

	return []int{}
}
