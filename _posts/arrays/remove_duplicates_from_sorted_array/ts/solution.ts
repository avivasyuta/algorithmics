export const removeDuplicates = (nums: number[]): number => {
    let pivot = 0

    for (let i = 0; i < nums.length - 1; i++) {
        if (nums[i] != nums[i + 1]) {
            nums[pivot] = nums[i]
            pivot++
        }
    }

    return pivot + 1
}