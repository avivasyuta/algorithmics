package number_of_islands

func numIslands(grid [][]byte) int {
	counter := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				counter++
				deepScan(grid, i, j)
			}
		}
	}

	return counter
}

func deepScan(grid [][]byte, i, j int) {
	grid[i][j] = '0'
	if j < len(grid[i])-1 && grid[i][j+1] == '1' {
		deepScan(grid, i, j+1)
	}
	if i < len(grid)-1 && grid[i+1][j] == '1' {
		deepScan(grid, i+1, j)
	}
	if j > 0 && grid[i][j-1] == '1' {
		deepScan(grid, i, j-1)
	}
	if i > 0 && grid[i-1][j] == '1' {
		deepScan(grid, i-1, j)
	}
}
