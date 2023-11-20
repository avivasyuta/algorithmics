package main

func generateParenthesisBruteforce(n int) []string {
	combinations := generateAllPossibleParenthesis(n)
	res := make([]string, 0, len(combinations))
	for _, parenthesis := range combinations {
		if !isValid(parenthesis) {
			continue
		}

		res = append(res, parenthesis)
	}

	return res
}

func generateAllPossibleParenthesis(n int) []string {
	combinations := make([]string, 0, 100)
	combinations = append(combinations, "(")

	for i := 1; i < n*2; i++ {
		newArray := make([]string, 0, len(combinations))
		for _, item := range combinations {
			newArray = append(newArray, item+"(")
			newArray = append(newArray, item+")")
		}

		combinations = newArray
	}

	return combinations
}

func isValid(brackets string) bool {
	stack := make([]rune, 0, len(brackets))

	closeBracket := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range brackets {
		if openBracket, isClose := closeBracket[char]; isClose {
			if len(stack) == 0 {
				return false
			}

			lastOpenBracket := stack[len(stack)-1]

			if lastOpenBracket != openBracket {
				return false
			}

			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}
