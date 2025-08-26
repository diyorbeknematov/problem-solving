package arraysstrings_test

import (
	arraysstrings "leetcode/arrays-strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{
			nums:     []int{1, 2, 3, 4},
			expected: []int{24, 12, 8, 6},
		},
		{
			nums:     []int{-1, 1, 0, -3, 3},
			expected: []int{0, 0, 9, 0, 0},
		},
	}

	for i, tt := range tests {
		res := arraysstrings.ProductExceptSelf(tt.nums)

		assert.Equal(t, tt.expected, res, "test case %d failed", i+1)
	}
}
