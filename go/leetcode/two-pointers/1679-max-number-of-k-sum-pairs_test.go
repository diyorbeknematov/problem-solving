package twopointers_test

import (
	twopointers "leetcode/two-pointers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxOperations(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected int
	}{
		{
			nums:     []int{1, 2, 3, 4},
			k:        5,
			expected: 2,
		},
		{
			nums:     []int{3, 1, 3, 4, 3},
			k:        6,
			expected: 1,
		},
	}

	for i, tt := range tests {
		res := twopointers.MaxOperations(tt.nums, tt.k)

		assert.Equal(t, tt.expected, res, "test case %d failed", i+1)
	}
}
