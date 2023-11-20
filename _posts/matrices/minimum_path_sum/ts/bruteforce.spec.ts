import { minPathSumBruteforce } from "./bruteforce";

type Suit = {
    name: string
    grid: number[][]
    expected: number
}

const suits: Suit[] = [
    {
        name: "example_1",
        grid: [
            [1, 3, 1],
            [1, 5, 1],
            [4, 2, 1],
        ],
        expected: 7,
    },
    {
        name: "example_2",
        grid: [
            [1, 2, 3],
            [4, 5, 6],
        ],
        expected: 12,
    },
    {
        name: "example_3",
        grid: [
            [1, 3, 1],
            [1, 5, 1],
            [1, 9, 1],
        ],
        expected: 7,
    },
    {
        name: "example_4",
        grid: [
            [1],
            [1],
            [1],
        ],
        expected: 3,
    },
    {
        name: "example_5",
        grid: [
            [1, 1, 2],
        ],
        expected: 4,
    },
]

describe('minPathSumBruteforce', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const res = minPathSumBruteforce(suit.grid)
            expect(res).toEqual(suit.expected);
        })
    })
})
