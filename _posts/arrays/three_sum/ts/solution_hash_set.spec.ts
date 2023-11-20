import { threeSumHashSet } from "./solution_hash_set";

type Suit = {
    name: string
    nums: number[]
    expRes: number[][]
}

const suits: Suit[] = [
    {
        name: "[-2, 0, 0, 2, 2]",
        nums: [-2, 0, 0, 2, 2],
        expRes: [[-2, 2, 0]],
    },
    {
        name: "[-1, 0, 1, 2, -1, -4]",
        nums: [-1, 0, 1, 2, -1, -4],
        expRes: [[-1, 1, 0], [-1, 2, -1]],
    },
    {
        name: "[0, 1, 1]",
        nums: [0, 1, 1],
        expRes: [],
    },
    {
        name: "[0, 0, 0]",
        nums: [0, 0, 0],
        expRes: [[0, 0, 0]],
    },
]

describe('removeDuplicates', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const res = threeSumHashSet(suit.nums)
            expect(res).toEqual(suit.expRes);
        })
    })
})

