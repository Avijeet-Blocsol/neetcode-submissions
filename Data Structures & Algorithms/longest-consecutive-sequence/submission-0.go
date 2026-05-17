func longestConsecutive(nums []int) int {
	check := make(map[int]bool)

	for _, i := range nums {
		check[i] = true
	}

	largest := 0

	for _, i := range nums {

		seqLen := 0

		if ok := check[i-1]; !ok {
			num := i

			for check[num] {
				num += 1
				seqLen += 1
			}

			if seqLen > largest {
				largest = seqLen
			}
		}
	}

	return largest
}
