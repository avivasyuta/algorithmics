export const numIslands = (grid: string[][]): number => {
    let counter = 0

    for (let i = 0; i < grid.length; i++) {
        for (let j = 0; j < grid[i].length; j++) {
            if (grid[i][j] == '1') {
                counter++
                deepScan(grid, i, j)
            }
        }
    }

    return counter
}

export const deepScan = (grid: string[][], i: number, j: number): void => {
    grid[i][j] = '0'

    if (j < grid[i].length - 1 && grid[i][j + 1] == '1') {
        deepScan(grid, i, j + 1)
    }
    if (i < grid.length - 1 && grid[i + 1][j] == '1') {
        deepScan(grid, i + 1, j)
    }
    if (j > 0 && grid[i][j - 1] == '1') {
        deepScan(grid, i, j - 1)
    }
    if (i > 0 && grid[i - 1][j] == '1') {
        deepScan(grid, i - 1, j)
    }
}
