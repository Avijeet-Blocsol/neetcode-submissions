func maxArea(heights []int) int {
max := 0

	l := 0
	r := len(heights) - 1

	for l < r {
		lp := heights[l]
		rp := heights[r]

		if lp < rp {
			sm := lp

			a := sm * (r - l)

			if a > max {
				max = a
			}

			l += 1
		} else {
			sm := rp

			a := sm * (r - l)

			if a > max {
				max = a
			}

			r -= 1
		}
	}

	return max
}
