func sortColors(nums []int) {
    check := [3]int{}

	for _, num := range nums {
		check[num] += 1
	}

	index := 0

	for index < check[0] {
		nums[index] = 0
		index += 1
	}

	for index < check[0]+check[1] {
		nums[index] = 1
		index += 1
	}

	for index < check[0]+check[1]+check[2] {
		nums[index] = 2
		index += 1
	}
}
