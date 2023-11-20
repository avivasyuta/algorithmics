package main

func twoSum(nums []int, target int) []int {
	// Создаем мапу для отслеживания индексом чисел из слайса.
	indexMap := make(map[int]int)

	for i, num := range nums {
		// Рассчитываем второе искомое число
		complement := target - num

		// Проверяем, есть ли индекс для этого числа в мапе
		secondIndex, ok := indexMap[complement]

		// Если индекс для числа уже есть в мапе, значит мы нашли искомую пару
		if ok {
			return []int{i, secondIndex}
		}

		// Если индекс для числа еще не добавлен, то записываем его в мапу
		indexMap[num] = i
	}

	return nil
}
