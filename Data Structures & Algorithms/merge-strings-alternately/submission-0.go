func mergeAlternately(word1 string, word2 string) string {

	var res strings.Builder

	a := 0
	b := 0

	for a < len(word1) && b < len(word2) {
		if a <= b {
			res.WriteString(string(word1[a]))
			a += 1
		} else {
			res.WriteString(string(word2[b]))
			b += 1
		}
	}

	if a < len(word1) {
		res.WriteString(word1[a:])
	}

	if b < len(word2) {
		res.WriteString(word2[b:])
	}

	return res.String()
}
