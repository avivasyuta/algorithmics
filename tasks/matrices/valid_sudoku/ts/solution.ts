export const isValidSudoku = (board: string[][]): boolean => {
	const n = board.length

	const rows: boolean[][] = []
	const columns: boolean[][] = []
	const quads: Record<string, boolean[]> = {}

	// Предварительно заполняем массивы false значениями
	for (let i = 0; i < n; i++) {
		rows.push(new Array(n).fill(false))
		columns.push(new Array(n).fill(false))
	}

	for (let row = 0; row < n; row++) {
		for (let col = 0; col < n; col++) {
			if (board[row][col] === '.') {
				continue
			}

			const pos = Number(board[row][col]) - 1
			const key = `${Math.floor(row / 3)}:${Math.floor(col / 3)}`

			if (!quads[key]) {
				quads[key] = new Array(n).fill(false)
			}

			if (rows[row][pos] || columns[col][pos] || quads[key][pos]) {
				return false
			}

			rows[row][pos] = true
			columns[col][pos] = true
			quads[key][pos] = true
		}
	}

	return true
}
