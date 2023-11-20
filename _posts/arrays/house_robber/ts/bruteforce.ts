type Seq = {
	index: number
	sum:   number
}

export function robBruteforce(nums: number[]): number {
	const sequences = genAllSequences(nums)
	let max = 0

	debugger;

	sequences.forEach((sequence) => {
		const lastIndex = sequence.length - 1
		if (sequence[lastIndex].sum > max) {
			max = sequence[lastIndex].sum
		}
	})

	return max
}

function genAllSequences(nums: number[]): Seq[][] {
	switch (nums.length) {
	case 0:
		return [[
			{
				index: 0,
				sum: 0,
			},
		]];
	case 1:
		return [[
			{
				index: 0,
				sum: nums[0],
			},
		]];
	case 2:
		return [[
			{
				index: 0,
				sum: Math.max(nums[0], nums[1]),
			},
		]];
	}

	let res: Seq[][] = [
		[
			{
				index: 0,
				sum:   nums[0],
			},
		],
		[
			{
				index: 1,
				sum:   nums[1],
			},
		],
		[
			{
				index: 0,
				sum:   nums[0],
			},
			{
				index: 2,
				sum:   nums[2] + nums[0],
			},
		],
	]

	for (let i = 3; i < nums.length; i++) {
		let temp: Seq[][] = []

		res.forEach((sequence) => {
			const last = sequence[sequence.length - 1]

			if (last.index === i - 2 || last.index === i - 3) {
				const tmpSeq = [
					...sequence,
					{
						index: i,
						sum:   last.sum + nums[i],
					}
				]

				temp.push(tmpSeq)
			} else {
				temp.push(sequence)
			}
		})

		res = temp
	}

	return res
}
