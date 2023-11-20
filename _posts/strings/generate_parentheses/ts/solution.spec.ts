import { generateParenthesis } from "./solution";
import { generateParenthesisBruteforce } from "./solution_bruteforce";

type Suit = {
	name: string
	in:   number
	out:  string[]
}

const suits: Suit[] = [
	{
		name: "3",
		in:   3,
		out:  ["((()))", "(()())", "(())()", "()(())", "()()()"],
	},
	{
		name: "1",
		in:   1,
		out:  ["()"],
	},
]

describe('generateParenthesis', () => {
	suits.forEach(suit => {
		test(suit.name, () => {
			const res = generateParenthesis(suit.in)
			expect(res).toEqual(suit.out);
		})
	})
})

describe('generateParenthesisBruteforce', () => {
	suits.forEach(suit => {
		test(suit.name, () => {
			const res = generateParenthesisBruteforce(suit.in)
			expect(res).toEqual(suit.out);
		})
	})
})