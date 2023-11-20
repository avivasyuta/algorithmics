import { numIslands } from "./solution";

type suit = {
    name: string
    grid: string[][]
    expected: number
}

const suits: suit[] = [
    {
        name: "example_1",
        grid: [
            ['1', '1', '1', '1', '0'],
            ['1', '1', '0', '1', '0'],
            ['1', '1', '0', '0', '0'],
            ['0', '0', '0', '0', '0'],
        ],
        expected: 1,
    },
    {
        name: "example_2",
        grid: [
            ['1', '1', '0', '0', '0'],
            ['1', '1', '0', '0', '0'],
            ['0', '0', '1', '0', '0'],
            ['0', '0', '0', '1', '1'],
        ],
        expected: 3,
    },
]

describe('numIslands', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const res = numIslands(suit.grid)
            expect(res).toEqual(suit.expected);
        })
    })
})
