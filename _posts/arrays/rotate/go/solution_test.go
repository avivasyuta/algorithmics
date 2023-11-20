package rotate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Suit struct {
	name   string
	nums   []int
	k      int
	expRes []int
}

func Test_removeDuplicates(t *testing.T) {
	suits := []Suit{
		{
			name:   "nums = [1, 2, 3, 4, 5, 6, 7], k = 3",
			nums:   []int{1, 2, 3, 4, 5, 6, 7},
			k:      3,
			expRes: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name:   "nums = [-1,-100,3,99], k = 2",
			nums:   []int{-1, -100, 3, 99},
			k:      2,
			expRes: []int{3, 99, -1, -100},
		},
		{
			name:   "nums = [1,2,3,4], k = 6",
			nums:   []int{1, 2, 3, 4},
			k:      6,
			expRes: []int{3, 4, 1, 2},
		},
		{
			name:   "nums = [], k = 6",
			nums:   []int{},
			k:      6,
			expRes: []int{},
		},
		{
			name:   "nums = [1], k = 2",
			nums:   []int{1},
			k:      2,
			expRes: []int{1},
		},
		{
			name:   "nums = [1, 2, 3], k = 0",
			nums:   []int{1, 2, 3},
			k:      0,
			expRes: []int{1, 2, 3},
		},
	}

	for _, suit := range suits {
		t.Run(suit.name, func(t *testing.T) {
			rotate(suit.nums, suit.k)
			assert.Equal(t, suit.expRes, suit.nums)
		})
	}
}
