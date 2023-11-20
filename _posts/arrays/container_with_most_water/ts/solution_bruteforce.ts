export const maxAreaBruteforce = (height: number[]): number => {
    let res = 0

    for (let i = 0; i < height.length - 1; i++) {
        for (let j = height.length - 1; j > i; j--) {
            const length = j - i

            if (length * height[i] < res) {
                break
            }

            let minHeight = height[i]

            if (height[j] < minHeight) {
                minHeight = height[j]
            }

            const square = minHeight * length

            if (square > res) {
                res = square
            }
        }
    }

    return res
}
