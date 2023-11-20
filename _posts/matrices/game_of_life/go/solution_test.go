package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type suitGameOfLife struct {
	name     string
	grid     [][]int
	expected [][]int
}

func Test_gameOfLife(t *testing.T) {
	testCases := []suitGameOfLife{
		{
			name: "example_1",
			grid: [][]int{
				{0, 1, 0},
				{0, 0, 1},
				{1, 1, 1},
				{0, 0, 0},
			},
			expected: [][]int{
				{0, 0, 0},
				{1, 0, 1},
				{0, 1, 1},
				{0, 1, 0},
			},
		},
		{
			name: "example_2",
			grid: [][]int{
				{1, 1},
				{1, 0},
			},
			expected: [][]int{
				{1, 1},
				{1, 1},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			gameOfLife(testCase.grid)
			assert.Equal(t, testCase.expected, testCase.grid)
		})
	}
}
