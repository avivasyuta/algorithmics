import { longestCommonPrefix } from "./solution";

type Suit = {
    name: string
    in: string[]
    expected: string
}

const suits: Suit[] = [
    {
        name: "[flower, flow, flight]",
        in: ["flower", "flow", "flight"],
        expected: "fl",
    },
    {
        name: "[dog, racecar, car]",
        in: ["dog", "racecar", "car"],
        expected: "",
    },
    {
        name: "[dog, dogs]",
        in: ["dog", "dogs"],
        expected: "dog",
    },
    {
        name: "[dog, dogs, ]",
        in: ["dog", "dogs", ""],
        expected: "",
    },
    {
        name: "[dog]",
        in: ["dog"],
        expected: "dog",
    },
]

describe('longestCommonPrefix', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const res = longestCommonPrefix(suit.in)
            expect(res).toEqual(suit.expected);
        })
    })
})
