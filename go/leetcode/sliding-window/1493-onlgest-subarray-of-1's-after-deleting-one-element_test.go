package slidingwindow_test

import (
	slidingwindow "leetcode/sliding-window"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestSubarray(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{1, 1, 0, 1},
			expected: 3,
		},
		{
			nums:     []int{0, 1, 1, 1, 0, 1, 1, 0, 1},
			expected: 5,
		},
		{
			nums:     []int{1, 1, 1},
			expected: 2,
		},
	}

	for i, tt := range tests {
		res := slidingwindow.LongestSubarray(tt.nums)

		assert.Equal(t, tt.expected, res, "test case %d failed", i+1)
	}
}
