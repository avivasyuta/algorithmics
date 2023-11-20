import { myPow } from "./solution";
import { myPowBruteforce } from "./bruteforce";

type Suit = {
    name: string
    x: number
    n: number
    out: number
}

const suits: Suit[] = [
    {
        name: "2^10",
        x: 2.0,
        n: 10,
        out: 1024,
    },
    {
        name: "2.1^3",
        x: 2.1,
        n: 3,
        out: 9.26100,
    },
    {
        name: "2^-2",
        x: 2,
        n: -2,
        out: 0.25,
    },
    {
        name: "1^123456",
        x: 1,
        n: 123456,
        out: 1,
    },
    {
        name: "-1^123456",
        x: -1,
        n: 123456,
        out: 1,
    },
    {
        name: "-1^123457",
        x: -1,
        n: 123457,
        out: -1,
    },
]

describe('myPowBruteforce', () => {
	suits.forEach(suit => {
		test(suit.name, () => {
			const res = myPowBruteforce(suit.x, suit.n)
			expect(res).toEqual(suit.out);
		})
	})
})

describe('myPow', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const res = myPow(suit.x, suit.n)
            expect(res).toEqual(suit.out);
        })
    })
})