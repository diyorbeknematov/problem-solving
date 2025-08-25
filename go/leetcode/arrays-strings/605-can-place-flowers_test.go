package arraysstrings_test

import (
	arraysstrings "leetcode/arrays-strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanPlaceFlowers(t *testing.T) {
	tests := []struct {
		flowerbed []int
		n         int
		expected  bool
	}{
		{
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         1,
			expected:  true,
		},
		{
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         2,
			expected:  false,
		},
		{
			flowerbed: []int{1, 0, 0, 0, 0, 1},
			n:         2,
			expected:  false,
		},
		{
			flowerbed: []int{0, 0, 1, 0, 0},
			n:         1,
			expected:  true,
		},
	}

	for i, tt := range tests {
		res := arraysstrings.CanPlaceFlowers(tt.flowerbed, tt.n)

		assert.Equal(t, tt.expected, res, "Test case %d failed", i+1)
	}
}
