func sortArray(nums []int) []int {
return mergeSort(nums)
}

func mergeSort(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	mid := len(nums) / 2

	a := mergeSort(nums[:mid])
	b := mergeSort(nums[mid:])

	c := merge(a, b)

	return c
}

func merge(a, b []int) []int {
	n := len(a)
	p := len(b)

	res := []int{}

	r := 0
	t := 0

	for r < n && t < p {
		if a[r] < b[t] {
			res = append(res, a[r])
			r++
		} else {
			res = append(res, b[t])
			t++
		}
	}

	if r < n {
		res = append(res, a[r:]...)
	}

	if t < p {
		res = append(res, b[t:]...)
	}

	return res
}

