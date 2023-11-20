export const reverse = (nums: number[], left: number, right: number): void => {
    while (left < right) {
        const tmp = nums[left]
        nums[left] = nums[right]
        nums[right] = tmp
        left++
        right--
    }
}

export const rotate = (nums: number[], k: number): void => {
    const length = nums.length
    if (length < 2) {
        return
    }

    const count = k % length

    reverse(nums, 0, length-1)
    reverse(nums, 0, count-1)
    reverse(nums, count, length-1)
};
