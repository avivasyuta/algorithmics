package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type twoSumSuit struct {
	name   string
	nums   []int
	target int
	expRes []int
}

func Test_twoSum(t *testing.T) {
	suits := []twoSumSuit{
		{
			name:   "[2, 7, 11, 15]",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			expRes: []int{1, 2},
		},
		{
			name:   "[2, 3, 4]",
			nums:   []int{2, 3, 4},
			target: 6,
			expRes: []int{1, 3},
		},
		{
			name:   "[-1, 0]",
			nums:   []int{-1, 0},
			target: -1,
			expRes: []int{1, 2},
		},
		{
			name:   "[0, 0]",
			nums:   []int{0, 0},
			target: 5,
			expRes: []int{-1, -1},
		},
	}

	for _, suit := range suits {
		t.Run(suit.name, func(t *testing.T) {
			res := twoSum(suit.nums, suit.target)
			assert.Equal(t, suit.expRes, res)
		})
	}
}
