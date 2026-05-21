func removeDuplicates(nums []int) int {
i := 0
	j := 0
	u := 0

	for j < len(nums) {

		c := nums[j]

		if j+1 < len(nums) && nums[j+1] == c {
			for j+1 < len(nums) && nums[j+1] == c {
				j += 1
			}

			nums[i] = c
			i += 1
			u += 1
			j += 1
		} else {
			nums[i] = nums[j]
			i += 1
			j += 1
			u += 1
		}
	}

	return u
}
