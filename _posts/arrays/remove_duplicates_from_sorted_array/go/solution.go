package main

func removeDuplicates(nums []int) int {
	pivot := 0

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			nums[pivot] = nums[i]
			pivot++
		}
	}

	return pivot + 1
}
