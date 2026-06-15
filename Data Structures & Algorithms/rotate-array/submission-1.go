func rotate(nums []int, k int) {
	n := len(nums)

	if k > n {
		k = k % n
	}

	l := 0
	r := n - 1

	for l < r {
		nums[l], nums[r] = nums[r], nums[l]

		l += 1
		r -= 1
	}

	l = 0
	r = k - 1

	for l < r {
		nums[l], nums[r] = nums[r], nums[l]

		l += 1
		r -= 1
	}

	l = k
	r = n - 1

	for l < r {
		nums[l], nums[r] = nums[r], nums[l]

		l += 1
		r -= 1
	}
}
