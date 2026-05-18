func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	
	var builder strings.Builder

	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			builder.WriteRune(r)
		}
	}

	cleanStr := builder.String()
	n := len(cleanStr)

	for i := range n {
		if cleanStr[i] != cleanStr[n-1-i] {
			return false
		}
	}

	return true
}
