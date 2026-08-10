func sortColors(nums []int) {
	i := 0
	c := 0
	j := len(nums) - 1

	for c <= j && j > 0 {
		if nums[c] == 0 {
			nums[i], nums[c] = nums[c], nums[i]
			c++
			i++
			continue
		}

		if nums[c] == 1 {
			c++
			continue
		}

		if nums[c] == 2 {
			nums[c], nums[j] = nums[j], nums[c]
			j--
		}
	}
}
