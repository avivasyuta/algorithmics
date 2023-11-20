import { twoSum } from "./solution";

type Suit = {
    name: string
    nums: number[]
    target: number
    expRes: number[]
}

const suits: Suit[] = [
    {
        name: "[2, 7, 11, 15], 9",
        nums: [2, 7, 11, 15],
        target: 9,
        expRes: [1, 0],
    },
    {
        name: "[3, 2, 4], 6",
        nums: [3, 2, 4],
        target: 6,
        expRes: [2, 1],
    },
    {
        name: "[-1, 0], -1",
        nums: [-1, 0],
        target: -1,
        expRes: [1, 0],
    },
    {
        name: "[0, 0], 5",
        nums: [0, 0],
        target: 5,
        expRes: [],
    },
]

describe('twoSum', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const res = twoSum(suit.nums, suit.target)
            expect(res).toEqual(suit.expRes);
        })
    })
})