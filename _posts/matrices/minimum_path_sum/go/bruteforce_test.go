package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minPathSumBruteforce(t *testing.T) {
	testCases := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name: "example_1",
			grid: [][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			},
			expected: 7,
		},
		{
			name: "example_2",
			grid: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			expected: 12,
		},
		{
			name: "example_3",
			grid: [][]int{
				{1, 3, 1},
				{1, 5, 1},
				{1, 9, 1},
			},
			expected: 7,
		},
		{
			name: "example_4",
			grid: [][]int{
				{1},
				{2},
				{3},
			},
			expected: 6,
		},
		{
			name: "example_5",
			grid: [][]int{
				{1, 2, 3},
			},
			expected: 6,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := minPathSumBruteforce(testCase.grid)
			assert.Equal(t, testCase.expected, res)
		})
	}
}
