package slidingwindow_test

import (
	slidingwindow "leetcode/sliding-window"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaxAverage(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected float64
	}{
		{
			nums:     []int{1, 12, -5, -6, 50, 3},
			k:        4,
			expected: 12.75000,
		},
		{
			nums:     []int{5},
			k:        1,
			expected: 5.00000,
		},
	}

	for i, tt := range tests {
		res := slidingwindow.FindMaxAverage(tt.nums, tt.k)

		assert.Equal(t, tt.expected, res, "test case %d failed", i+1)
	}
}
