func sortColors(nums []int) {
	check := [3]int{}

	for _, num := range nums {
		check[num] += 1
	}

	index := 0

	for i, count := range check {
		for count > 0 {
			nums[index] = i
			count -= 1
			index += 1
		}
	}
}
