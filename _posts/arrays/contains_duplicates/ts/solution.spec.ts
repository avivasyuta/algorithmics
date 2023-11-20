import { containsDuplicateWithMap } from "./solution_map";
import { containsDuplicateWithSort } from "./solution_sort";

type Suit = {
    name: string
    nums: number[]
    expRes: boolean
}

const suits: Suit[] = [
    {
        name: "[1, 2, 3]",
        nums: [1, 2, 3],
        expRes: false,
    },
    {
        name: "[2, 3, 4, 2]",
        nums: [2, 3, 4, 2],
        expRes: true,
    },
    {
        name: "[]",
        nums: [],
        expRes: false,
    },
]

describe('containsDuplicate with map', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const res = containsDuplicateWithMap(suit.nums)
            expect(res).toEqual(suit.expRes);
        })
    })
})

describe('containsDuplicate wit sort', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const res = containsDuplicateWithSort(suit.nums)
            expect(res).toEqual(suit.expRes);
        })
    })
})
