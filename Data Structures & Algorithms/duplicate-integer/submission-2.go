func hasDuplicate(nums []int) bool {
    seen := make(map[int]bool)

	for i := range nums {
		if _, exists := seen[nums[i]]; exists {
			return true
		}

		seen[nums[i]] = true
	}

	return false
}
