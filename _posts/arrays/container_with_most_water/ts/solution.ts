export const maxArea = (height: number[]): number => {
    let res = 0
    let i = 0
    let j = height.length - 1

    while (i < j) {
        const length = j - i
        let minHeight = height[i]

        if (height[j] < minHeight) {
            minHeight = height[j]
        }

        const square = length * minHeight

        if (square > res) {
            res = square
        }

        if (height[i] < height[j]) {
            i++
            continue
        }

        if (height[j] < height[i]) {
            j--
            continue
        }

        i++
        j--
    }

    return res
}
