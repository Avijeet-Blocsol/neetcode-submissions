import "slices"

func threeSum(nums []int) [][]int {
	c := make(map[[3]int]bool, 0)

	slices.Sort(nums)

	// fmt.Printf("nums: %v \n", nums)

	for idx, num := range nums {

		// fmt.Printf("num: %v \n", num)

		cp := TwoSumModified(idx, 0-num, nums)

		for _, pair := range cp {
			r := []int{num}
			r = append(r, pair...)

			// fmt.Printf("r is %v \n", r)

			if len(r) != 3 {
				continue
			}

			slices.Sort(r)

			var key [3]int
			copy(key[:], r)
			if ok := c[key]; !ok {
				c[key] = true
			}
		}
	}

	result := make([][]int, 0, len(c))

	for k, _ := range c {
		result = append(result, k[:])
	}

	return result
}

func TwoSumModified(index, target int, nums []int) [][]int {

	// fmt.Printf("index : %v, target : %v \n", index, target)

	l := 0
	r := len(nums) - 1

	ret := make([][]int, 0)

	for l < r {

		// fmt.Printf("l : %v, r : %v \n", l, r)

		if l == index {
			l += 1
			continue
		}

		if r == index {
			r -= 1
			continue
		}

		res := nums[l] + nums[r]

		if res == target {
			ret = append(ret, []int{nums[l], nums[r]})
			l += 1
			r -= 1
			continue
		}

		if res < target {
			l += 1
		} else {
			r -= 1
		}
	}

	return ret
}
