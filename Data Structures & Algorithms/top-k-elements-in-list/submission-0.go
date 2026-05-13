func topKFrequent(nums []int, k int) []int {
	check := map[int]int{}

	for _, num := range nums {
		if val, ok := check[num]; ok {
			check[num] = val + 1
		} else {
			check[num] = 1
		}
	}

	freq_arr := make([][]int, len(nums)+1)
	result := []int{}

	for key, val := range check {
		freq_arr[val] = append(freq_arr[val], key)
	}

	for i := len(freq_arr) - 1; i > 0; i-- {

		for j := 0; j < len(freq_arr[i]); j++ {
			if k == 0 {
				break
			}
			result = append(result, freq_arr[i][j])
			k -= 1
		}
		if k == 0 {
			break
		}
	}

	return result
}
