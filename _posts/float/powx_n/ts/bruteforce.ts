const powAbsNBruteforce = (x: number, n: number): number => {
    if (n === 0) {
        return 1
    }

    if (n === 1) {
        return x
    }

    let res = x

    for (n; n > 1; n--) {
        res *= x
    }

    return res
}

export const myPowBruteforce = (x: number, n: number): number => {
    if (x === 1) {
        return 1
    }

    if (x == -1) {
        if (n % 2 === 0) {
            return 1
        }

        return -1
    }

    let res: number

    if (n < 0) {
        res = 1 / powAbsNBruteforce(x, -1 * n)
    } else {
        res = powAbsNBruteforce(x, n)
    }

    return parseFloat(res.toFixed(5))
}
