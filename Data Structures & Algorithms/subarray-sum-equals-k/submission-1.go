func subarraySum(nums []int, k int) int {
	total := 0
	curr := 0

	prefix := make(map[int]int)

	for _, num := range nums {
		prefix[curr] += 1
		curr += num

		total += prefix[curr-k]
	}

	return total
}
