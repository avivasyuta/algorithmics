package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type removeDuplicatesSuit struct {
	name   string
	nums   []int
	expRes int
}

func Test_removeDuplicates(t *testing.T) {
	suits := []removeDuplicatesSuit{
		{
			name:   "[0,0,1,1,1,2,2,3,3,4]",
			nums:   []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expRes: 5,
		},
		{
			name:   "[1,1,2]",
			nums:   []int{1, 1, 2},
			expRes: 2,
		},
		{
			name:   "[1,1,1,1]",
			nums:   []int{1, 1, 1, 1},
			expRes: 1,
		},
		{
			name:   "[1,2,3,4]",
			nums:   []int{1, 2, 3, 4},
			expRes: 4,
		},
		{
			name:   "[1,2,2,2]",
			nums:   []int{1, 2, 2, 2},
			expRes: 2,
		},
		{
			name:   "[1]",
			nums:   []int{1},
			expRes: 1,
		},
	}

	for _, suit := range suits {
		t.Run(suit.name, func(t *testing.T) {
			res := removeDuplicates(suit.nums)
			assert.Equal(t, suit.expRes, res)
		})
	}
}
