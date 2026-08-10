
import "slices"

type Freq struct {
	Val  int
	Freq int
}


func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)

	for i := range nums {
		freqMap[nums[i]]++
	}

	vals := make([]Freq, 0, len(freqMap))

	for num, count := range freqMap {
		vals = append(vals, Freq{Val: num, Freq: count})
	}

	slices.SortFunc(vals, func(a, b Freq) int {
		return b.Freq - a.Freq
	})

	result := make([]int, 0, k)
	
	for i := 0; i < k && i < len(vals); i++ {
		result = append(result, vals[i].Val)
	}

	return result
}
