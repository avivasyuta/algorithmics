package main

func maxArea(height []int) int {
	res := 0
	i := 0
	j := len(height) - 1

	for i < j {
		length := j - i
		minHeight := height[i]
		if height[j] < minHeight {
			minHeight = height[j]
		}

		square := length * minHeight
		if square > res {
			res = square
		}

		if height[i] < height[j] {
			i++
			continue
		}

		if height[j] < height[i] {
			j--
			continue
		}

		i++
		j--
	}

	return res
}
