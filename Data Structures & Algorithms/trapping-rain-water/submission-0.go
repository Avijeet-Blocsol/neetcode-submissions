func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}

	left := 0
	right := len(height) - 1

	left_max := 0
	right_max := 0

	t := 0

	for left < right {
		if height[left] <= height[right] {
			if height[left] > left_max {
				left_max = height[left]
				continue
			} else {
				t += left_max - height[left]
				left += 1
			}
		} else {
			if height[right] > right_max {
				right_max = height[right]
				continue
			} else {
				t += right_max - height[right]
				right -= 1
			}
		}
	}

	return t
}
