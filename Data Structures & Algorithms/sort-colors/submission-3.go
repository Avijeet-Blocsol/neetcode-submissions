func sortColors(nums []int) {
    a := 0
	b := 0
	c := 0

	for i := range nums {
		switch nums[i] {
		case 0:
			a++
		case 1:
			b++
		default:
			c++
		}
	}

	for i := range a {
		nums[i] = 0
	}

	for i := a; i < a+b; i++ {
		nums[i] = 1
	}

	for i := a + b; i < a+b+c; i++ {
		nums[i] = 2
	}
}
