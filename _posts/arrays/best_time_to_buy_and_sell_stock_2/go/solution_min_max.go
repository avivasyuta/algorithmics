package main

func maxProfitMinMax(prices []int) int {
	res := 0
	min := prices[0]
	max := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] > max {
			max = prices[i]
		} else {
			res += max - min
			min = prices[i]
			max = prices[i]
		}
	}

	res += max - min

	return res
}
