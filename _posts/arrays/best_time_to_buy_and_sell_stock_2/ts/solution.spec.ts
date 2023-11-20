import { maxProfitMinMax } from "./solution_min_max";
import { maxProfitTrends } from "./solution_trends";

type RemoveDuplicatesSuit = {
    name: string
    prices: number[]
    expRes: number
}

const suits: RemoveDuplicatesSuit[] = [
    {
        name: "[7,1,5,3,6,4]",
        prices: [7, 1, 5, 3, 6, 4],
        expRes: 7,
    },
    {
        name: "[1,2,3,4,5,6]",
        prices: [1, 2, 3, 4, 5, 6],
        expRes: 5,
    },
    {
        name: "[6,5,3,1]",
        prices: [6, 5, 3, 1],
        expRes: 0,
    },
]

describe('maxProfitMinMax', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const res = maxProfitMinMax(suit.prices)
            expect(res).toEqual(suit.expRes);
        })
    })
})

describe('maxProfitTrends', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const res = maxProfitTrends(suit.prices)
            expect(res).toEqual(suit.expRes);
        })
    })
})