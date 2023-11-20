export const longestCommonPrefix = (strs: string[]): string => {
	if (strs.length == 0) {
		return ""
	}

	let prefix = ''
	const firstStr = strs[0]

	for (let i = 0; i < firstStr.length; i++) {
		const char = firstStr[i]

		for (let j = 0; j < strs.length; j++) {
			const str = strs[j]
			if (i >= str.length || str[i] != firstStr[i]) {
				return prefix
			}
		}

		prefix += char
	}

	return prefix
}
