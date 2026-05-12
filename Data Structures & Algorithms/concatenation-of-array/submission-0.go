func getConcatenation(nums []int) []int {
    n := len(nums)

	result := make([]int, n*2)

	for i := 0; i < n*2; i++ {
		if i < n {
			result[i] = nums[i]
		} else {
			result[i] = nums[i-n]
		}
	}

	return result
}
