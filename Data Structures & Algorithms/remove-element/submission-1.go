import "slices"

func removeElement(nums []int, val int) int {
    count := 0

	for i := range nums {
		if nums[i] == val {
			nums[i] = 51
		} else {
			count += 1
		}
	}

	slices.Sort(nums)

	return count
}
