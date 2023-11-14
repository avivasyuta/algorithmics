package valid_sudoku

import (
	"fmt"
	"strconv"
)

func isValidSudoku(board [][]byte) bool {
	rows := [9][9]bool{}
	columns := [9][9]bool{}
	quads := make(map[string][9]bool, 9)

	n := len(board)

	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			if board[row][col] == '.' {
				continue
			}

			strNum := string(board[row][col])
			num, _ := strconv.Atoi(strNum)
			pos := num - 1

			key := fmt.Sprintf("%d:%d", row/3, col/3)

			if rows[row][pos] || columns[col][pos] || quads[key][pos] {
				return false
			}

			rows[row][pos] = true
			columns[col][pos] = true

			temp := quads[key]
			temp[pos] = true
			quads[key] = temp
		}
	}

	return true
}
