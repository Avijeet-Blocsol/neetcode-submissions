func containsNearbyDuplicate(nums []int, k int) bool {

	m := map[int]int{}

	for idx, num := range nums {
		if v, ok := m[num]; ok {
			if idx-v <= k {
				return true
			} 
			m[num] = idx
		} else {
			m[num] = idx
		}
	}

	return false
}
