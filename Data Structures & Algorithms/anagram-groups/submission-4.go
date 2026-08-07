func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string)

	for i := range strs {
		c := strs[i]

		g := [26]int{}

		for j := range len(c) {
			g[c[j]-'a'] += 1
		}

		groups[g] = append(groups[g], c)
	}

	res := [][]string{}

	for _, v := range groups {
		res = append(res, v)
	}

	return res
}
