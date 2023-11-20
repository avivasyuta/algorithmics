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
			name:   "[2, 7, 11, 15], 9",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			expRes: []int{1, 0},
		},
		{
			name:   "[3, 2, 4], 6",
			nums:   []int{3, 2, 4},
			target: 6,
			expRes: []int{2, 1},
		},
		{
			name:   "[-1, 0], -1",
			nums:   []int{-1, 0},
			target: -1,
			expRes: []int{1, 0},
		},
		{
			name:   "[0, 0], 5",
			nums:   []int{0, 0},
			target: 5,
			expRes: nil,
		},
	}

	for _, suit := range suits {
		t.Run(suit.name, func(t *testing.T) {
			res := twoSum(suit.nums, suit.target)
			assert.Equal(t, suit.expRes, res)
		})
	}
}
