package number_of_islands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type suit struct {
	name     string
	grid     [][]byte
	expected int
}

func Test_minPathSum(t *testing.T) {
	testCases := []suit{
		{
			name: "example_1",
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			expected: 1,
		},
		{
			name: "example_2",
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			expected: 3,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := numIslands(testCase.grid)
			assert.Equal(t, testCase.expected, res)
		})
	}
}
