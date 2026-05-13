func sortArray(nums []int) []int {
	return mergeSort(nums)
}

func mergeSort(nums []int) []int {

	if len(nums) == 1 {
		return nums
	}

	left := mergeSort(nums[:len(nums)/2])
	right := mergeSort(nums[len(nums)/2:])

	return merge(left, right)
}

func merge(a, b []int) []int {
	a1 := 0
	b1 := 0

	res := []int{}

	for a1 < len(a) && b1 < len(b) {
		if a[a1] < b[b1] {
			res = append(res, a[a1])
			a1 += 1
		} else {
			res = append(res, b[b1])
			b1 += 1
		}
	}

	if a1 < len(a) {
		res = append(res, a[a1:]...)
	}

	if b1 < len(b) {
		res = append(res, b[b1:]...)
	}

	return res
}
