func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	target := make(map[byte]int)

	for i := range t {
		target[t[i]]++
	}

	form := 0
	min := ""
	l := 0

	check := make(map[byte]int)

	for r := range len(s) {
		char := s[r]
		check[char]++

		if _, exists := target[char]; exists && target[char] == check[char] {
			form++
		}

		for l <= r && form == len(target) {
			// We increase l till window becomes invalid
			char_l := s[l]

			if min == "" || r-l+1 < len(min) {
				min = s[l : r+1]
			}

			check[char_l]--

			if _, exists := target[char_l]; exists && check[char_l] < target[char_l] {
				form--
			}

			l++
		}

	}

	return min
}
