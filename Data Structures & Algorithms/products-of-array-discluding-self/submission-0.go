func productExceptSelf(nums []int) []int {
	n := len(nums)
	leftProduct := make([]int, n)
	rightProduct := make([]int, n)
	result := make([]int, n)

	for i := range n {
		leftProduct[i] = 1
		rightProduct[i] = 1
	}

	current := 1

	for i := 1; i < n; i++ {
		current *= nums[i-1]
		leftProduct[i] = current
	}

	current = 1

	for i := n - 2; i >= 0; i-- {
		current *= nums[i+1]
		rightProduct[i] = current
	}

	for i := range n {
		result[i] = leftProduct[i] * rightProduct[i]
	}

	return result
}
