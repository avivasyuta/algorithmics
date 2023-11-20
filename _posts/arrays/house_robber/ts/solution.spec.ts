import { robBruteforce } from "./bruteforce";
import { rob } from "./solution";

type Case = {
	name:     string
	in:       number[]
	expected: number
}

const testCases: Case[] = [
	{
		name: "[1,2,3,1]",
		in: [1, 2, 3, 1],
		expected: 4,
	},
	{
		name: "[2,7,9,3,1]",
		in: [2, 7, 9, 3, 1],
		expected: 12,
	},
	{
		name: "[1]",
		in: [1],
		expected: 1,
	},
	{
		name: "[2,1]",
		in: [2, 1],
		expected: 2,
	},
]

describe('robBruteforce', () => {
	testCases.forEach(suit => {
		test(suit.name, () => {
			const res = robBruteforce(suit.in)
			expect(res).toEqual(suit.expected);
		})
	})
})

describe('rob', () => {
	testCases.forEach(suit => {
		test(suit.name, () => {
			const res = rob(suit.in)
			expect(res).toEqual(suit.expected);
		})
	})
})
