export const twoSum = (nums: number[], target: number): number[] => {
    // Устанавливаем левый указатель в нулевой индекс
    let left = 0

    // Устанавливаем правый указатель в индекс последнего элемента в массиве
    let right = nums.length - 1

    // Запускаем цикл, который прервется, если `left` и `right` сойдутся
    while (left < right) {
        const sum = nums[left] + nums[right]

        if (sum === target) {
            return [left + 1, right + 1]
        } else if (sum > target) {
            right--
        } else {
            left++
        }
    }

    // Если ничего не нашли возвращаем [-1, -1]
    return [-1, -1]
}
