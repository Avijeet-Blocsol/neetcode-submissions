func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	freq := make([]int, 26)

	for i := range s {
		freq[s[i]-'a'] += 1
		freq[t[i]-'a'] -= 1
	}

	for i := range freq {
		if freq[i] != 0 {
			return false
		}
	}

	return true
}
