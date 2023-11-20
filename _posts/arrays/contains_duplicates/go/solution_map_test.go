package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type WithMapSuit struct {
	name   string
	nums   []int
	expRes bool
}

func Test_containsDuplicate_withMap(t *testing.T) {
	suits := []WithMapSuit{
		{
			name:   "[1, 2, 3]",
			nums:   []int{1, 2, 3},
			expRes: false,
		},
		{
			name:   "[2, 3, 4, 2]",
			nums:   []int{2, 3, 4, 2},
			expRes: true,
		},
		{
			name:   "[]",
			nums:   []int{},
			expRes: false,
		},
	}

	for _, suit := range suits {
		t.Run(suit.name, func(t *testing.T) {
			res := containsDuplicateWithMap(suit.nums)
			assert.Equal(t, suit.expRes, res)
		})
	}
}
