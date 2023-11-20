package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestCommonPrefix(t *testing.T) {
	testCases := []struct {
		name     string
		in       []string
		expected string
	}{
		{
			name:     "[flower, flow, flight]",
			in:       []string{"flower", "flow", "flight"},
			expected: "fl",
		},
		{
			name:     "[dog, racecar, car]",
			in:       []string{"dog", "racecar", "car"},
			expected: "",
		},
		{
			name:     "nil",
			in:       nil,
			expected: "",
		},
		{
			name:     "[dog, dogs]",
			in:       []string{"dog", "dogs"},
			expected: "dog",
		},
		{
			name:     "[dog, dogs, ]",
			in:       []string{"dog", "dogs", ""},
			expected: "",
		},
		{
			name:     "[dog]",
			in:       []string{"dog"},
			expected: "dog",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := longestCommonPrefix(testCase.in)
			assert.Equal(t, testCase.expected, res)
		})
	}
}
