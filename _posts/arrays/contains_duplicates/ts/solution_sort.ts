export const containsDuplicateWithSort = (nums: number[]): boolean => {
	nums.sort()

	for (let i = 0; i < nums.length - 1; i++) {
		if (nums[i] === nums[i + 1]) {
			return true
		}
	}

	return false
}
