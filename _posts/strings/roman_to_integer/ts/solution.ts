const runeToIntegerMap: Record<string, number> = {
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

const runeToIntegerDecrementsMap: Record<string, Record<string, number>> = {
	'I': {
		'V': 4,
		'X': 9,
	},
	'X': {
		'L': 40,
		'C': 90,
	},
	'C': {
		'D': 400,
		'M': 900,
	},
}

export const romanToInt = (s: string): number => {
	let res = 0

	for (let i = 0; i < s.length; i++) {
		if (i < s.length - 1) {
			const possibleDecrements = runeToIntegerDecrementsMap[s[i]]

			if (possibleDecrements) {
				const realNum = possibleDecrements[s[i + 1]]

				if (realNum) {
					res += realNum
					i++
					continue
				}
			}
		}

		res += runeToIntegerMap[s[i]]
	}

	return res
}
