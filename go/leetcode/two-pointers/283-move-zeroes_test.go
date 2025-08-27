package twopointers_test

import (
	twopointers "leetcode/two-pointers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{
			nums:     []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			nums:     []int{0},
			expected: []int{0},
		},
	}

	for i, tt := range tests {
		twopointers.MoveZeroes(tt.nums)

		assert.Equal(t, tt.expected, tt.nums, "test case %d failed", i+1)
	}
}
