func firstMissingPositive(nums []int) int {
	n := len(nums)

	for i, num := range nums {
		if num < 0 || num > n {
			nums[i] = 1<<31 - 1
		}
	}

	for i, num := range nums {
		if num != 1<<31-1 && num > 0 {
			n := num
			curr := nums[num-1]

			for n != 1<<31-1 && curr != n {
				nums[n-1] = n
				nums[i] = curr
				n = curr

				if n != 1<<31-1 && n > 0 {
					curr = nums[n-1]
				}
			}
		}
	}

	for i, num := range nums {
		if num == 1<<31-1 || num != i+1 {
			return i + 1
		}
	}
	return n + 1
}
