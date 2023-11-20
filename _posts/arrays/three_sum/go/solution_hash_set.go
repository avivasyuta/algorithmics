package main

import "sort"

func threeSumHashSet(nums []int) [][]int {
	// Создаем слайс для результатов
	result := make([][]int, 0)

	// Предварительно сортируем слайс по возрастанию
	sort.Ints(nums)

	for i, firstNum := range nums {
		// Не проверяем текущее число, если оно такое же, как и предыдущее, потому для него мы получим такой же результат.
		if i != 0 && nums[i-1] == firstNum {
			continue
		}

		indexMap := make(map[int]bool)

		for j := i + 1; j < len(nums); j++ {
			secondNum := nums[j]

			// Высчитываем третье искомое число
			thirdNum := 0 - firstNum - secondNum

			_, ok := indexMap[thirdNum]
			if ok {
				result = append(result, []int{firstNum, secondNum, thirdNum})
				for j+1 < len(nums) && nums[j] == nums[j+1] {
					j++
				}
			}

			indexMap[secondNum] = true
		}
	}

	return result
}
