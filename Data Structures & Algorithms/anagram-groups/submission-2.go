func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string)

	for _, val := range strs {
		rep := make([]int, 26)

		for _, runic := range val {
			rep[runic-'a'] += 1
		}

		if _, ok := groups[[26]int(rep)]; ok {
			groups[[26]int(rep)] = append(groups[[26]int(rep)], val)
		} else {
			groups[[26]int(rep)] = []string{val}
		}
	}

	result := make([][]string, 0)

	for _, val := range groups {
		result = append(result, val)
	}

	return result
}
