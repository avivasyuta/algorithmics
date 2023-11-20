package main

type seq struct {
	index int
	sum   int
}

func robBruteforce(nums []int) int {
	sequences := genAllSequences(nums)
	max := 0

	for _, sequence := range sequences {
		if sequence[len(sequence)-1].sum > max {
			max = sequence[len(sequence)-1].sum
		}
	}

	return max
}

func genAllSequences(nums []int) [][]seq {
	switch len(nums) {
	case 0:
		return [][]seq{
			{
				{sum: 0},
			},
		}
	case 1:
		return [][]seq{
			{
				{sum: nums[0]},
			},
		}
	case 2:
		max := nums[0]
		if nums[1] > max {
			max = nums[1]
		}

		return [][]seq{
			{
				{sum: max},
			},
		}
	}

	res := [][]seq{
		{
			{
				index: 0,
				sum:   nums[0],
			},
		},
		{
			{
				index: 1,
				sum:   nums[1],
			},
		},
		{
			{
				index: 0,
				sum:   nums[0],
			},
			{
				index: 2,
				sum:   nums[2] + nums[0],
			},
		},
	}

	for i := 3; i < len(nums); i++ {
		temp := make([][]seq, 0, len(res))
		for _, sequence := range res {
			last := sequence[len(sequence)-1]

			if last.index == i-2 || last.index == i-3 {
				tmpSeq := make([]seq, 0, len(sequence))
				tmpSeq = append(sequence, seq{
					index: i,
					sum:   last.sum + nums[i],
				})
				temp = append(temp, tmpSeq)
			} else {
				temp = append(temp, sequence)
			}
		}

		res = temp
	}

	return res
}
