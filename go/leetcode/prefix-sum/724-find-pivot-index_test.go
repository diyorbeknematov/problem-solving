package prefixsum_test

import (
	prefixsum "leetcode/prefix-sum"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPivotIndex(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{1, 7, 3, 6, 5, 6},
			expected: 3,
		},
		{
			nums:     []int{1, 2, 3},
			expected: -1,
		},
		{
			nums:     []int{2, 1, -1},
			expected: 0,
		},
	}

	for i, tt := range tests {
		res := prefixsum.PivotIndex(tt.nums)

		assert.Equal(t, tt.expected, res, "test case %d failed", i+1)
	}
}
