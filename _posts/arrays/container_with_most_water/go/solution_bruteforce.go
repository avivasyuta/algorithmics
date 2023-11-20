package main

func maxAreaBruteforce(height []int) int {
	res := 0
	for i := 0; i < len(height)-1; i++ {
		for j := len(height) - 1; j > i; j-- {
			length := j - i
			if length*height[i] < res {
				break
			}

			minHeight := height[i]
			if height[j] < minHeight {
				minHeight = height[j]
			}

			square := minHeight * length
			if square > res {
				res = square
			}
		}
	}

	return res
}
