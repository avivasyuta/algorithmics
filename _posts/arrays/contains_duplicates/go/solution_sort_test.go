package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type WithSortSuit struct {
	name   string
	nums   []int
	expRes bool
}

func Test_containsDuplicate_withSort(t *testing.T) {
	suits := []WithSortSuit{
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
			res := containsDuplicateWithSort(suit.nums)
			assert.Equal(t, suit.expRes, res)
		})
	}
}
