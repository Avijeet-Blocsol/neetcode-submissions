func hasDuplicate(nums []int) bool {
    c := map[int]bool{}

	for _, val := range nums {
		if !c[val] {
			c[val] = true
		} else {
			return true
		}
	}

	return false
}
