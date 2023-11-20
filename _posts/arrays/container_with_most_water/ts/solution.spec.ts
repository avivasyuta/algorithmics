import { maxArea } from "./solution";
import { maxAreaBruteforce } from "./solution_bruteforce";

type Case = {
	name:     string
	in:       number[]
	expected: number
}

const testCases: Case[] = [
	{
		name:     "[1, 8, 6, 2, 5, 4, 8, 3, 7]",
		in:       [1, 8, 6, 2, 5, 4, 8, 3, 7],
		expected: 49,
	},
	{
		name:     "[1, 1]",
		in:       [1, 1],
		expected: 1,
	},
	{
		name:     "[0, 0]",
		in:       [0, 0],
		expected: 0,
	},
	{
		name:     "[0, 1]",
		in:       [0, 1],
		expected: 0,
	},
	{
		name:     "[1, 8, 6, 2, 5, 4, 8, 3, 7]",
		in:       [1, 8, 6, 2, 5, 4, 8, 3, 7],
		expected: 49,
	},
	{
		name:     "[2, 3, 4, 5, 18, 17, 6]",
		in:       [2, 3, 4, 5, 18, 17, 6],
		expected: 17,
	},
]

describe('maxArea', () => {
	testCases.forEach(suit => {
		test(suit.name, () => {
			const res = maxArea(suit.in)
			expect(res).toEqual(suit.expected);
		})
	})
})

describe('maxAreaBruteforce', () => {
	testCases.forEach(suit => {
		test(suit.name, () => {
			const res = maxAreaBruteforce(suit.in)
			expect(res).toEqual(suit.expected);
		})
	})
})

