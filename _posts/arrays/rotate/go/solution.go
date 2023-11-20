package rotate

func reverse(nums []int, left int, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

func rotate(nums []int, k int) {
	length := len(nums)
	if length < 2 {
		return
	}

	count := k % length

	reverse(nums, 0, length-1)
	reverse(nums, 0, count-1)
	reverse(nums, count, length-1)
}
