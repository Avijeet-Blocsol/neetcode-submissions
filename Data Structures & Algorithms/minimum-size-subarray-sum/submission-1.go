func minSubArrayLen(target int, nums []int) int {
	min := len(nums) + 1
	curr := 0
	l := 0
	r := 0

	for r < len(nums) {
		curr += nums[r]

		if curr >= target {

			for curr >= target {
				if r-l+1 < min {
					min = r - l + 1
				}

				curr -= nums[l]

				l += 1
			}
		}

		r += 1
	}

	if min == len(nums)+1 {
		return 0
	}

	return min
}
