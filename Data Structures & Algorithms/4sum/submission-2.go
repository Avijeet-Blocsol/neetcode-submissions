
import "slices"

func fourSum(nums []int, target int) [][]int {

	c := make(map[[4]int]bool, 0)

	slices.Sort(nums)

	for idx, num := range nums {

		cp := ThreeSumModifed(idx, target-num, nums)

		for _, threeSum := range cp {
			r := [4]int{num}

			for i := 1; i < 4; i++ {
				r[i] = threeSum[i-1]
			}

			if len(r) != 4 {
				continue
			}

			slices.SortFunc(r[:], func(a, b int) int {
				return a - b
			})

			if ok := c[r]; !ok {
				c[r] = true
			}
		}
	}

	result := make([][]int, 0, len(c))

	for k := range c {
		result = append(result, k[:])
	}

	return result
}

func ThreeSumModifed(index, target int, nums []int) [][]int {

	c := make(map[[3]int]bool, 0)

	for idx, num := range nums {

		if idx == index {
			continue
		}

		cp := TwoSumModified2(index, idx, target-num, nums)

		for _, pair := range cp {
			r := []int{num}
			r = append(r, pair...)

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

	for k := range c {
		result = append(result, k[:])
	}

	return result
}

func TwoSumModified2(index1, index2, target int, nums []int) [][]int {

	l := 0
	r := len(nums) - 1

	ret := make([][]int, 0)

	for l < r {

		if l == index1 || l == index2 {
			l += 1
			continue
		}

		if r == index1 || r == index2 {
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
