func removeElement(nums []int, val int) int {
    unique := 0

	for i := range nums {
		if nums[i] == val {
			nums[i] = 51
		} else {
			unique++
		}
	}

	idx := len(nums) - 1

	for idx >= 0 {
		if nums[idx] == 51 {
			idx--
		} else {
			break
		}
	}

	for i := idx - 1; i >= 0; i-- {
		if nums[i] == 51 {
			nums[i], nums[idx] = nums[idx], nums[i]
			idx--
		}
	}

	return unique
}
