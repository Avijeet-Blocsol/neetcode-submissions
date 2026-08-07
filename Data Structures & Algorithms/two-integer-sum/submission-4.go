func twoSum(nums []int, target int) []int {
	res := []int{}

	check := map[int]int{}

	for i := range nums {
		if v, exists := check[target-nums[i]]; exists {
			res = []int{v, i}
			break
		}

		check[nums[i]] = i
	}

	return res
}
