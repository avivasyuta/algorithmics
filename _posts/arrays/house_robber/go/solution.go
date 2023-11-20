package main

func rob(nums []int) int {
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	case 2:
		max := nums[0]
		if nums[1] > max {
			max = nums[1]
		}

		return max
	}

	nums[2] = nums[2] + nums[0]

	for i := 3; i < len(nums); i++ {
		max := nums[i-2]
		if nums[i-3] > max {
			max = nums[i-3]
		}

		nums[i] = max + nums[i]
	}

	max := nums[len(nums)-1]
	if nums[len(nums)-2] > max {
		max = nums[len(nums)-2]
	}

	return max
}
