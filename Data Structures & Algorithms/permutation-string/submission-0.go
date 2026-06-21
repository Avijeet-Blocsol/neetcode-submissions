func checkInclusion(s1 string, s2 string) bool {
	n1 := len(s1)
	n2 := len(s2)

	if n1 > n2 {
		return false
	}

	o := make([]int, 26)

	for _, char := range s1 {
		o[char-'a'] += 1
	}

	c := make([]int, 26)

	for i := range n1 {
		c[s2[i]-'a'] += 1
	}

	g := func(a, b []int) bool {
		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}

		return true
	}

	if g(o, c) {
		return true
	}

	for i := n1; i < n2; i++ {
		c[s2[i-n1]-'a'] -= 1
		c[s2[i]-'a'] += 1

		if g(o, c) {
			return true
		}
	}

	return false
}
