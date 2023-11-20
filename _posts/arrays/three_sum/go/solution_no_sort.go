package main

import (
	"sort"
	"strconv"
	"strings"
)

func threeSumNoSort(nums []int) [][]int {
	uniq := make(map[string]bool)
	dups := make(map[int]bool)
	seen := make(map[int]int)

	for i, firstNum := range nums {
		_, ok := dups[firstNum]
		if ok {
			continue
		}

		dups[firstNum] = true

		for j := i + 1; j < len(nums); j++ {
			secondNum := nums[j]

			// Высчитываем третье искомое число
			thirdNum := 0 - firstNum - secondNum

			idx, ok := seen[thirdNum]

			if ok && idx == i {
				// Собираем уникальную тройку
				triplet := []string{strconv.Itoa(firstNum), strconv.Itoa(secondNum), strconv.Itoa(thirdNum)}

				// Сортируем тройку, чтобы избежать дублей
				sort.Strings(triplet)

				// Конвертируем тройку в строку и добавляем в map
				str := strings.Join(triplet, ".")
				uniq[str] = true
			}

			seen[secondNum] = i
		}
	}

	// Разбираем map с уникальными ключами обратно в слайс троек
	res := make([][]int, 0, len(uniq))

	for key := range uniq {
		numStrs := strings.Split(key, ".")
		resNums := make([]int, 0, len(numStrs))

		for _, numStr := range numStrs {
			num, _ := strconv.Atoi(numStr)
			resNums = append(resNums, num)
		}
		res = append(res, resNums)
	}

	return res
}
