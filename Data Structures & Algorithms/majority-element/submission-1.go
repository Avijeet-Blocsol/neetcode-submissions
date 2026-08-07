func majorityElement(nums []int) int {
	m := 1000000001
	freq := 0

	for i := range nums {
		if nums[i] != m {
			if freq > 0 {
				freq--
				continue
			} else {
				m = nums[i]
				freq = 1
			}
		} else {
			freq += 1
		}
	}

	return m
}
