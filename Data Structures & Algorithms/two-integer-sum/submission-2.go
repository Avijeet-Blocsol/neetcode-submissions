func twoSum(nums []int, target int) []int {
	check_map := map[int]int{}

	for i, num := range nums {
		if check_map[target-num] != 0 {
			if i < check_map[target-num]-1 {
				return []int{i, check_map[target-num] - 1}
			}

			return []int{check_map[target-num] - 1, i}
		} else {
			check_map[num] = i + 1
		}
	}

	return []int{}
}
