package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type threeSumTwoPointersSuit struct {
	name   string
	nums   []int
	expRes [][]int
}

func Test_threeSumTwoPointers(t *testing.T) {
	suits := []threeSumTwoPointersSuit{
		{
			name:   "[-2, 0, 0, 2, 2]",
			nums:   []int{-2, 0, 0, 2, 2},
			expRes: [][]int{{-2, 0, 2}},
		},
		{
			name:   "[-1, 0, 1, 2, -1, -4]",
			nums:   []int{-1, 0, 1, 2, -1, -4},
			expRes: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name:   "[0, 1, 1]",
			nums:   []int{0, 1, 1},
			expRes: [][]int{},
		},
		{
			name:   "[0, 0, 0]",
			nums:   []int{0, 0, 0},
			expRes: [][]int{{0, 0, 0}},
		},
	}

	for _, suit := range suits {
		t.Run(suit.name, func(t *testing.T) {
			res := threeSumTwoPointers(suit.nums)
			assert.Equal(t, suit.expRes, res)
		})
	}
}
