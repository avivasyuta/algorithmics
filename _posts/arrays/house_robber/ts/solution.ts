export function rob(nums: number[]): number {
    switch (nums.length) {
        case 0:
            return 0
        case 1:
            return nums[0]
        case 2:
            return Math.max(nums[0], nums[1])
    }

    nums[2] = nums[2] + nums[0]

    for (let i = 3; i < nums.length; i++) {
        const max = Math.max(nums[i - 2], nums[i - 3])
        nums[i] = max + nums[i]
    }

    return Math.max(nums[nums.length - 1], nums[nums.length - 2])
}
