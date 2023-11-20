import { isValid } from "./solution";

type Suit = {
	name: string
	in: string
	out: boolean
}

const suits: Suit[] = [
	{
		name: "()",
		in:   "()",
		out:  true,
	},
	{
		name: "()[]{}",
		in:   "()[]{}",
		out:  true,
	},
	{
		name: "(]",
		in:   "(]",
		out:  false,
	},
	{
		name: "((((",
		in:   "((((",
		out:  false,
	},
	{
		name: "))))",
		in:   "))))",
		out:  false,
	},
	{
		name: "(){",
		in:   "(){",
		out:  false,
	},
	{
		name: "}{",
		in:   "}{",
		out:  false,
	},
]

describe('isValidBrackets', () => {
	suits.forEach(suit => {
		test(suit.name, () => {
			const res = isValid(suit.in)
			expect(res).toEqual(suit.out);
		})
	})
})
