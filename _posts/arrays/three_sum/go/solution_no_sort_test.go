package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type threeSumNoSortSuit struct {
	name   string
	nums   []int
	expRes [][]int
}

func Test_threeSumNoSort(t *testing.T) {
	suits := []threeSumNoSortSuit{
		{
			name:   "[-2, 0, 0, 2, 2]",
			nums:   []int{-2, 0, 0, 2, 2},
			expRes: [][]int{{-2, 0, 2}},
		},
		{
			name:   "[-1, 0, 1, 2, -1, -4]",
			nums:   []int{-1, 0, 1, 2, -1, -4},
			expRes: [][]int{{-1, 0, 1}, {-1, -1, 2}},
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
			res := threeSumNoSort(suit.nums)
			assert.Equal(t, suit.expRes, res)
		})
	}
}
