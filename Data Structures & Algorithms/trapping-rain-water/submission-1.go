func trap(height []int) int {
		if len(height) <= 2 {
		return 0
	}

	left := 0
	right := len(height) - 1

	left_max := height[left]
	right_max := height[right]

	t := 0

	for left < right {
		if height[left] <= height[right] {
			left++
			if height[left] > left_max {
				left_max = height[left]
			} else {
				t += left_max - height[left]
			}
		} else {
			right--
			if height[right] > right_max {
				right_max = height[right]
			} else {
				t += right_max - height[right]
			}
		}
	}

	return t
}
