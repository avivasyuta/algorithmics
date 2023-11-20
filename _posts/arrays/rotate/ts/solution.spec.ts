import { rotate } from "./solution";

type Suit = {
    name: string
    nums: number[]
    k: number
    expRes: number[]
}

const suits: Suit[] = [
    {
        name: "nums = [1, 2, 3, 4, 5, 6, 7], k = 3",
        nums: [1, 2, 3, 4, 5, 6, 7],
        k: 3,
        expRes: [5, 6, 7, 1, 2, 3, 4],
    },
    {
        name: "nums = [-1,-100,3,99], k = 2",
        nums: [-1, -100, 3, 99],
        k: 2,
        expRes: [3, 99, -1, -100],
    },
    {
        name: "nums = [1,2,3,4], k = 6",
        nums: [1, 2, 3, 4],
        k: 6,
        expRes: [3, 4, 1, 2],
    },
    {
        name: "nums = [], k = 6",
        nums: [],
        k: 6,
        expRes: [],
    },
    {
        name: "nums = [1], k = 2",
        nums: [1],
        k: 2,
        expRes: [1],
    },
    {
        name: "nums = [1, 2, 3], k = 0",
        nums: [1, 2, 3],
        k: 0,
        expRes: [1, 2, 3],
    },
]

describe('removeDuplicates', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            rotate(suit.nums, suit.k)
            expect(suit.nums).toEqual(suit.expRes);
        })
    })
})
