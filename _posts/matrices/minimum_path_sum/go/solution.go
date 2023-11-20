package main

func minPathSum(grid [][]int) int {
	var i, j int
	for i = 0; i < len(grid); i++ {
		for j = 0; j < len(grid[i]); j++ {
			if i == 0 && j == 0 {
				continue
			}

			if i == 0 {
				grid[i][j] = grid[i][j-1] + grid[i][j]

				continue
			}

			if j == 0 {
				grid[i][j] = grid[i-1][j] + grid[i][j]

				continue
			}

			minPrices := grid[i][j-1]
			if grid[i-1][j] < minPrices {
				minPrices = grid[i-1][j]
			}

			grid[i][j] = minPrices + grid[i][j]
		}
	}

	return grid[i-1][j-1]
}
