import { gameOfLife } from "./solution";

type Suit = {
    name: string
    grid: number[][]
    expected: number[][]
}

const suits: Suit[] = [
    {
        name: "example_1",
        grid: [
            [0, 1, 0],
            [0, 0, 1],
            [1, 1, 1],
            [0, 0, 0],
        ],
        expected: [
            [0, 0, 0],
            [1, 0, 1],
            [0, 1, 1],
            [0, 1, 0],
        ],
    },
    {
        name: "example_2",
        grid: [
            [1, 1],
            [1, 0],
        ],
        expected: [
            [1, 1],
            [1, 1],
        ],
    },
]

describe('gameOfLife', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            gameOfLife(suit.grid)
            expect(suit.grid).toEqual(suit.expected);
        })
    })
})
