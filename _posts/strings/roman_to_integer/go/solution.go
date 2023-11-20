package main

var runeToIntegerMap = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

var runeToIntegerDecrementsMap = map[rune]map[rune]int{
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

func romanToInt(s string) int {
	res := 0
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		if i < len(runes)-1 {
			if possibleDecrements, exist := runeToIntegerDecrementsMap[runes[i]]; exist {
				realNum, found := possibleDecrements[runes[i+1]]
				if found {
					res += realNum
					i++
					continue
				}
			}
		}

		realNum := runeToIntegerMap[runes[i]]
		res += realNum
	}

	return res
}
