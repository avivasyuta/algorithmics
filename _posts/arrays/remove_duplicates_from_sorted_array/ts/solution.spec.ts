import {removeDuplicates} from "./solution";

type RemoveDuplicatesSuit = {
    name: string
    nums: number[]
    expRes: number
}

const suits: RemoveDuplicatesSuit[] = [
    {
        name: "[0,0,1,1,1,2,2,3,3,4]",
        nums: [0, 0, 1, 1, 1, 2, 2, 3, 3, 4],
        expRes: 5,
    },
    {
        name: "[1,1,2]",
        nums: [1, 1, 2],
        expRes: 2,
    },
    {
        name: "[1,1,1,1]",
        nums: [1, 1, 1, 1],
        expRes: 1,
    },
    {
        name: "[1,2,3,4]",
        nums: [1, 2, 3, 4],
        expRes: 4,
    },
    {
        name: "[1,2,2,2]",
        nums: [1, 2, 2, 2],
        expRes: 2,
    },
    {
        name: "[1]",
        nums: [1],
        expRes: 1,
    },
]

describe('removeDuplicates', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const res = removeDuplicates(suit.nums)
            expect(res).toEqual(suit.expRes);
        })
    })
})
