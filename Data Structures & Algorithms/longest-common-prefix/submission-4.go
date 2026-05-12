func longestCommonPrefix(strs []string) string {
	result := ""

	for i := range strs[0] {
		for j := range strs {
			if i >= len(strs[j]) {
				return result
			}

			if strs[j][i] != strs[0][i] {
				return result
			}
		}

		result += string(strs[0][i])
	}

	return result
}
