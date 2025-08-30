package slidingwindow_test

import (
	slidingwindow "leetcode/sliding-window"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestOnes(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected int
	}{
		{
			nums:     []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
			k:        2,
			expected: 6,
		},
		{
			nums:     []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
			k:        3,
			expected: 10,
		},
	}

	for i, tt := range tests {
		res := slidingwindow.LongestOnes(tt.nums, tt.k)

		assert.Equal(t, tt.expected, res, "test case %d failed", i+1)
	}
}
