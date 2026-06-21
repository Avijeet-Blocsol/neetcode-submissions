func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	min := n + 1
	curr := 0
	l := 0
	r := 0

	for r < n {
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

	if min == n+1 {
		return 0
	}

	return min
}
