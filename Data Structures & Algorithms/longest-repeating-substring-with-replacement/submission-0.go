func characterReplacement(s string, k int) int {
	n := len(s)
	l := 0
	m := 1
	max := 0

	c := map[byte]int{}

	for i := range n {
		c[s[i]] += 1
		v := c[s[i]]

		if v > m {
			m = v
		}

		if i-l+1-m > k {
			c[s[l]] -= 1
			l += 1
		}

		
		if i-l+1 > max {
			max = i - l + 1
		}
	}

	return max
}
