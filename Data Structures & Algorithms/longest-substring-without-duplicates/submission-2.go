func lengthOfLongestSubstring(s string) int {
	
	n := len(s)

	if n <= 1 {
		return n
	}

	max := 0
	left := 0

	c := map[rune]int{}

	for idx, char := range s {
		if val, ok := c[char]; ok && val >= 0 {
			if idx-left > max {
				max = idx - left
			}

			// fmt.Printf("setting left: val is %v, char is %v \n", val, string(char))

			for i := left; i < val+1; i++ {
				c[rune(s[i])] = -1
			}

			left = val + 1
		}

		c[char] = idx

		// fmt.Printf("max is %v, left is %v, idx is %v \n", max, left, idx)

		if idx == n-1 {
			if idx-left+1 > max {
				max = idx - left + 1
			}
		}

	}

	return max
}
