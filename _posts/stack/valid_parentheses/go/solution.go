package main

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
