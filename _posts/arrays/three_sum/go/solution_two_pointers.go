package main

import "sort"

func threeSumTwoPointers(nums []int) [][]int {
	// Создаем слайс для результатов
	result := make([][]int, 0)

	// Предварительно сортируем слайс по возрастанию
	sort.Ints(nums)

	for i, num := range nums {
		// Не проверяем текущее число, если оно такое же, как и предыдущее, потому для него мы получим такой же результат.
		if i != 0 && nums[i-1] == num {
			continue
		}

		left := i + 1
		right := len(nums) - 1

		for left < right {
			sum := num + nums[left] + nums[right]

			if sum == 0 {
				result = append(result, []int{num, nums[left], nums[right]})
				right -= 1
				left += 1

				for left < right && nums[left] == nums[left-1] {
					left += 1
				}
			} else if sum > 0 {
				right -= 1
			} else {
				left = left + 1
			}
		}
	}

	return result
}
