func longestCommonPrefix(strs []string) string {
	n := len(strs)

	if n == 1 {
		return strs[0]
	}

	result := strs[0]

	for i := 1; i < n; i++ {
		curr := strs[i]

		check_len := min(len(curr), len(result))

		if check_len == 0 {
			return ""
		}

		j := 0

		for j < check_len {
			if result[j] != curr[j] {
				result = curr[0:j]

				j = check_len
			} else {
				j += 1

				if j == check_len && len(curr) < len(result) {
					result = curr
				}
			}
		}
	}

	return result
}
