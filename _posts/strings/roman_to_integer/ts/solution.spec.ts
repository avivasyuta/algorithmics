import { romanToInt } from "./solution";

type Suit = {
    name: string
    in: string
    expRes: number
}

const suits: Suit[] = [
    {
        name: "III",
        in: "III",
        expRes: 3,
    },
    {
        name: "IX",
        in: "IX",
        expRes: 9,
    },
    {
        name: "XL",
        in: "XL",
        expRes: 40,
    },
    {
        name: "LVIII",
        in: "LVIII",
        expRes: 58,
    },
    {
        name: "MCMXCIV",
        in: "MCMXCIV",
        expRes: 1994,
    },
]

describe('romanToInt', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const res = romanToInt(suit.in)
            expect(res).toEqual(suit.expRes);
        })
    })
})
