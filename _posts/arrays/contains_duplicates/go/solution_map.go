package main

func containsDuplicateWithMap(nums []int) bool {
	frequency := make(map[int]bool)

	for _, num := range nums {
		_, exists := frequency[num]
		if exists {
			return true
		}
		frequency[num] = true
	}

	return false
}
