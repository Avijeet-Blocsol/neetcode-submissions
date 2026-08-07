func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	check := map[byte]int{}

	for i := range s {
		check[s[i]] += 1
	}

	for j := range t {
		v, exists := check[t[j]]

		if !exists || v == 0 {
			return false
		}

		check[t[j]] -= 1
	}

	for _, val := range check {
		if val != 0 {
			return false
		}
	}

	return true
}
