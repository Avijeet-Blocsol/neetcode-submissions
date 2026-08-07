func longestCommonPrefix(strs []string) string {
    s := 200
	prefix := ""

	for i := range strs {
		if len(strs[i]) < s {
			s = len(strs[i])
		}
	}

	for j := range s {
		c := strs[0][j]

		for k := 1; k < len(strs); k++ {
			if strs[k][j] != c {
				return prefix
			}
		}

		prefix += string(c)
	}

	return prefix
}
