package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rob(t *testing.T) {
	testCases := []struct {
		name     string
		in       []int
		expected int
	}{
		{
			name: "[1,2,3,1]",
			in: []int{
				1, 2, 3, 1,
			},
			expected: 4,
		},
		{
			name: "[2,7,9,3,1]",
			in: []int{
				2, 7, 9, 3, 1,
			},
			expected: 12,
		},
		{
			name:     "nil",
			expected: 0,
		},
		{
			name: "[1]",
			in: []int{
				1,
			},
			expected: 1,
		},
		{
			name: "[2,1]",
			in: []int{
				2, 1,
			},
			expected: 2,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := rob(testCase.in)
			assert.Equal(t, testCase.expected, res)
		})
	}
}
