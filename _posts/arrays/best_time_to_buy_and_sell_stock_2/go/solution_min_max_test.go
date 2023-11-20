package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type maxProfitMinMaxSuit struct {
	name   string
	prices []int
	expRes int
}

func Test_maxProfitMinMax(t *testing.T) {
	var suits = []maxProfitMinMaxSuit{
		{
			name:   "[7,1,5,3,6,4]",
			prices: []int{7, 1, 5, 3, 6, 4},
			expRes: 7,
		},
		{
			name:   "[1,2,3,4,5,6]",
			prices: []int{1, 2, 3, 4, 5, 6},
			expRes: 5,
		},
		{
			name:   "[6,5,3,1]",
			prices: []int{6, 5, 3, 1},
			expRes: 0,
		},
	}

	for _, suit := range suits {
		t.Run(suit.name, func(t *testing.T) {
			res := maxProfitMinMax(suit.prices)
			assert.Equal(t, suit.expRes, res)
		})
	}
}
