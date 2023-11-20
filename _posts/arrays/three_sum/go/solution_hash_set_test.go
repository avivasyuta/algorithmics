package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type threeSumHashSetSuit struct {
	name   string
	nums   []int
	expRes [][]int
}

func Test_threeSumHashSet(t *testing.T) {
	suits := []threeSumHashSetSuit{
		{
			name:   "[-2, 0, 0, 2, 2]",
			nums:   []int{-2, 0, 0, 2, 2},
			expRes: [][]int{{-2, 2, 0}},
		},
		{
			name:   "[-1, 0, 1, 2, -1, -4]",
			nums:   []int{-1, 0, 1, 2, -1, -4},
			expRes: [][]int{{-1, 1, 0}, {-1, 2, -1}},
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
			res := threeSumHashSet(suit.nums)
			assert.Equal(t, suit.expRes, res)
		})
	}
}
