export const threeSumNoSort = (nums: number[]): number[][] => {
    const uniq = new Set<string>()
    const dups = new Set<number>()
    const seen = new Map<number, number>()

    for (let i = 0; i < nums.length; i++) {
        const firstNum = nums[i]

        if (dups.has(firstNum)) {
            continue
        }

        dups.add(firstNum)

        for (let j = i + 1; j < nums.length; j++) {
            const secondNum = nums[j]

            // Высчитываем третье искомое число
            const thirdNum = 0 - firstNum - secondNum

            if (seen.has(thirdNum) && seen.get(thirdNum) === firstNum) {
                // Собираем уникальную тройку
                const triplet = [firstNum, secondNum, thirdNum]

                // Сортируем тройку, чтобы избежать дублей
                triplet.sort()

                // Конвертируем тройку в строку и добавляем в set
                uniq.add(triplet.join('.'))
            }

            seen.set(secondNum, firstNum)
        }
    }

    // Конвертируем set строк в массив
    return Array.from(uniq).reduce<number[][]>((acc, tripletStr) => {
        const tripletArr = tripletStr.split('.')
        const triplet = tripletArr.map(numStr => Number(numStr))
        return [...acc, triplet]
    }, [])
}