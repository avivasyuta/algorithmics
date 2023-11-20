export const gameOfLife = (board: number[][]): void => {
    for (let i = 0; i < board.length; i++) {
        for (let j = 0; j < board[i].length; j++) {
            const count = liveCount(board, i, j)

            if (board[i][j] == 0) {
                if (count == 3) {
                    board[i][j] = 2
                }

                continue
            }

            if (count < 2 || count > 3) {
                board[i][j] = 3
                continue
            }

            board[i][j] = 1
        }
    }

    for (let i = 0; i < board.length; i++) {
        for (let j = 0; j < board[i].length; j++) {
            if (board[i][j] == 3) {
                board[i][j] = 0
            }

            if (board[i][j] == 2) {
                board[i][j] = 1
            }
        }
    }
}

function liveCount(board: number[][], i: number, j: number): number {
    const m = board[i].length
    const n = board.length
    
    let up = 0;
    let down = 0;
    let sides = 0;

    if (i > 0) {
        up = getValue(board[i - 1][j])

        if (j > 0) {
            up += getValue(board[i - 1][j - 1])
        }

        if (j < m - 1) {
            up += getValue(board[i - 1][j + 1])
        }
    }

    if (i < n - 1) {
        down = getValue(board[i + 1][j])

        if (j > 0) {
            up += getValue(board[i + 1][j - 1])
        }

        if (j < m - 1) {
            up += getValue(board[i + 1][j + 1])
        }
    }

    if (j > 0) {
        sides += getValue(board[i][j - 1])
    }

    if (j < m - 1) {
        sides += getValue(board[i][j + 1])
    }

    return up + down + sides
}

function getValue(val: number): number {
    if (val == 2) {
        return 0
    }

    if (val == 3) {
        return 1
    }

    return val
}