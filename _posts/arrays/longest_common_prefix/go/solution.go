package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := ""
	firstStr := strs[0]

	for idx, char := range firstStr {
		for _, str := range strs {
			if idx >= len(str) || str[idx] != firstStr[idx] {
				return prefix
			}
		}

		prefix += string(char)
	}

	return prefix
}
