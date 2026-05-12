func isAnagram(s string, t string) bool {
	check_string := make([]int, 26)

	for _, val := range s {
		check_string[val-'a'] += 1
	}

	for _, val := range t {
		check_string[val-'a'] -= 1
	}

	for _, val := range check_string {
		if val != 0 {
			return false
		}
	}

	return true
}
