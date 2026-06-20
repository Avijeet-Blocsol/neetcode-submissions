func maxProfit(prices []int) int {
	max := 0
	minPrice := 100

	for _, num := range prices {
		if num < minPrice {
			minPrice = num
		} else {
			if num-minPrice > max {
				max = num - minPrice
			}
		}
	}

	return max
}
