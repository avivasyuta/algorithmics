export const maxProfitMinMax = (prices: number[]): number => {
    let res = 0
    let min = prices[0]
    let max = prices[0]

    for (let i = 1; i < prices.length; i++) {
        if (prices[i] > max) {
            max = prices[i]
        } else {
            res += max - min
            min = prices[i]
            max = prices[i]
        }
    }

    res += max - min

    return res
}