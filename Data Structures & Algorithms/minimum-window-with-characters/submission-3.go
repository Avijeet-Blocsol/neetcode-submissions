func minWindow(s string, t string) string {
    if len(s) < len(t) {
		return ""
	}

	target := make(map[byte]int)

	for i := range t {
		target[t[i]]++
	}

	check := make(map[byte]int)

	l := 0
	min := ""
	correct := 0

	for r := range s {
		char_r := s[r]
		check[char_r] += 1

		if _, exists := target[char_r]; exists && target[char_r] == check[char_r] {
			correct += 1
		}

		for correct == len(target) && l <= r {
			char_l := s[l]

			if min == "" || r-l+1 < len(min) {
				min = s[l : r+1]
			}

			check[char_l] -= 1

			if _, exists := target[char_l]; exists && check[char_l] < target[char_l] {
				correct -= 1
			}

			l += 1
		}
	}

	return min
}
