type NumPos struct {
	Val int
	Idx int
}

type MaxHeapStack []NumPos

func (h MaxHeapStack) Len() int {
	return len(h)
}

func (h MaxHeapStack) Less(i, j int) bool {
	return h[i].Val > h[j].Val
}

func (h MaxHeapStack) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeapStack) Push(i any) {
	*h = append(*h, i.(NumPos))
}

func (h MaxHeapStack) Peek() int {
	return h[0].Val
}

func (h *MaxHeapStack) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}


func maxSlidingWindow(nums []int, k int) []int {
    result := []int{}

	if k <= 0 || len(nums) == 0 || len(nums) < k {
		return result
	}

	h := &MaxHeapStack{}

	heap.Init(h)

	for i := range k {
		heap.Push(h, NumPos{
			Val: nums[i],
			Idx: i,
		})
	}

	result = append(result, h.Peek())

	for i := k; i < len(nums); i++ {
		curr := nums[i]

		heap.Push(h, NumPos{
			Val: curr,
			Idx: i,
		})

		for (*h)[0].Idx < i-k+1 {
			heap.Pop(h)
		}

		result = append(result, (*h)[0].Val)

	}

	return result
}
