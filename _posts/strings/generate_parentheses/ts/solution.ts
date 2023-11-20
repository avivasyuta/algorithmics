export const generate = (n: number, row: string, opened: number, closed: number, answer: string[]): void => {
    if (row.length === n * 2) {
        answer.push(row)
        return
    }

    if (opened >= closed && opened < n) {
        generate(n, row + "(", opened + 1, closed, answer)
    }

    if (opened > closed) {
        generate(n, row + ")", opened, closed + 1, answer)
    }
}

export const generateParenthesis = (n: number): string[] => {
    const res: string[] = []
    generate(n, "(", 1, 0, res)
    return res
}
