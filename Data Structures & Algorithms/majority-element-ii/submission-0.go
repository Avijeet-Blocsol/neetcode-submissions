func majorityElement(nums []int) []int {
	result := []int{}

	cand_1, cand_2 := 0, 0

	vote_1 := 0
	vote_2 := 0

	for _, num := range nums {

		if num == cand_1 {
			vote_1 += 1
			continue
		}

		if num == cand_2 {
			vote_2 += 1
			continue
		}

		if vote_1 == 0 {
			cand_1 = num
			vote_1 = 1
			continue
		}

		if vote_2 == 0 {
			cand_2 = num
			vote_2 = 1
			continue
		}

		vote_1 -= 1
		vote_2 -= 1
	}

	if vote_1 > 0 {
		vote_1 = 0

		for _, n := range nums {
			if n == cand_1 {
				vote_1 += 1
			}
		}

		if vote_1 > len(nums)/3 {
			result = append(result, cand_1)
		}
	}

	if vote_2 > 0 {
		vote_2 = 0

		for _, n := range nums {
			if n == cand_2 {
				vote_2 += 1
			}
		}

		if vote_2 > len(nums)/3 {
			result = append(result, cand_2)
		}
	}

	return result
}
