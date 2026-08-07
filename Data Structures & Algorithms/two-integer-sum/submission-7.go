func twoSum(nums []int, target int) []int {
	check := map[int]int{}

	for i := range nums {
		if v, exists := check[target-nums[i]]; exists {
			return []int{v, i}
		}

		check[nums[i]] = i
	}

	return []int{}
}
