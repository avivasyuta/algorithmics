package main

func twoSum(nums []int, target int) []int {
	// Устанавливаем левый указатель в нулевой индекс
	left := 0

	// Устанавливаем правый указатель в индекс последнего элемента в слайсе
	right := len(nums) - 1

	// Запускаем цикл, который прервется, если `left` и `right` сойдутся
	for left < right {
		sum := nums[left] + nums[right]

		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum > target {
			right--
		} else {
			left++
		}
	}

	// Если ничего не нашли возвращаем [-1, -1]
	return []int{-1, -1}
}
