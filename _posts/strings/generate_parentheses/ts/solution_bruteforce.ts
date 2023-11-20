export function generateParenthesisBruteforce(n: number): string[] {
	const combinations = generateAllPossibleParenthesis(n)
	let res: string[] = []

	combinations.forEach((parenthesis) => {
		if (isValid(parenthesis)) {
			res.push(parenthesis)
		}
	})

	return res
}

function generateAllPossibleParenthesis(n: number): string[] {
	let combinations: string[] = ["("]

	for (let i = 1; i < n * 2; i++) {
		const newArray: string[] = []

		combinations.forEach((item) => {
			newArray.push(item + "(")
			newArray.push(item + ")")
		})

		combinations = newArray
	}

	return combinations
}

export const isValid = (brackets: string): boolean => {
	const stack: string[] = []

	const closeBrackets: Record<string, string> = {
		')': '(',
		']': '[',
		'}': '{',
	}

	for (let i = 0; i < brackets.length; i++) {
		const char = brackets[i]
		const openBracket = closeBrackets[char]

		if (openBracket) {
			if (stack.length === 0) {
				return false
			}

			const lastOpenBracket = stack[stack.length - 1]

			if (lastOpenBracket !== openBracket) {
				return false
			}

			stack.pop()
		} else {
			stack.push(char)
		}
	}

	return stack.length === 0
}

