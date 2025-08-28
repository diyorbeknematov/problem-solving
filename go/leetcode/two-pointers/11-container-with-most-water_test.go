package twopointers_test

import (
	twopointers "leetcode/two-pointers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T) {
	tests := []struct {
		height   []int
		expected int
	}{
		{
			height:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			height:   []int{1, 1},
			expected: 1,
		},
	}

	for i, tt := range tests {
		res := twopointers.MaxArea(tt.height)

		assert.Equal(t, tt.expected, res, "test case %d failed", i+1)
	}
}
