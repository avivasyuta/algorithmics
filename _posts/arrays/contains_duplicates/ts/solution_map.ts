export const containsDuplicateWithMap = (nums: number[]): boolean => {
	const frequency: Map<number, boolean> = new Map()

	for (let i = 0; i < nums.length; i++) {
		const num = nums[i]
		if (frequency.has(num)) {
			return true
		}
		frequency.set(num, true)
	}

	return false
}
