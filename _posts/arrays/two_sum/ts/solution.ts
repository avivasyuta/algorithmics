export const twoSum = (nums: number[], target: number): number[] => {
    // Создаем объект для отслеживания индексом чисел из массива
    // Также можно использовать new Map(), но объект делает в точности то же самое и с ним немного проще работаь.
    const indexMap: Record<number, number> = {}

    for (let i = 0; i < nums.length; i++) {
        // Рассчитываем второе искомое число
        const complement = target - nums[i]

        // Проверяем, есть ли индекс для этого числа в объекте
        const secondIndex = indexMap[complement]

        // Если индекс для числа уже есть в объекте, значит мы нашли искомую пару
        if (secondIndex > -1) {
            return [i, secondIndex]
        }

        // Если индекс для числа еще не добавлен, то записываем его в объект
        indexMap[nums[i]] = i
    }

    return []
}