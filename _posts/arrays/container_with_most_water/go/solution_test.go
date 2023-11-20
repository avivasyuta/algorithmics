package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	name     string
	in       []int
	expected int
}{
	{
		name:     "[1, 8, 6, 2, 5, 4, 8, 3, 7]",
		in:       []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
		expected: 49,
	},
	{
		name:     "[1, 1]",
		in:       []int{1, 1},
		expected: 1,
	},
	{
		name:     "[0, 0]",
		in:       []int{0, 0},
		expected: 0,
	},
	{
		name:     "[0, 1]",
		in:       []int{0, 1},
		expected: 0,
	},
	{
		name:     "[1, 8, 6, 2, 5, 4, 8, 3, 7]",
		in:       []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
		expected: 49,
	},
	{
		name:     "[2, 3, 4, 5, 18, 17, 6]",
		in:       []int{2, 3, 4, 5, 18, 17, 6},
		expected: 17,
	},
}

func Test_maxArea(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			res := maxArea(testCase.in)
			assert.Equal(t, testCase.expected, res)
		})
	}
}

func Test_maxAreaBruteforce(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			res := maxAreaBruteforce(testCase.in)
			assert.Equal(t, testCase.expected, res)
		})
	}
}
