package arraysstrings_test

import (
	arraysstrings "leetcode/arrays-strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncreasingTriplet(t *testing.T) {
	tests := []struct {
		nums     []int
		expected bool
	}{
		{
			nums:     []int{1, 2, 3, 4, 5},
			expected: true,
		},
		{
			nums:     []int{5, 4, 3, 2, 1},
			expected: false,
		},
		{
			nums:     []int{2, 1, 5, 0, 4, 6},
			expected: true,
		},
	}

	for i, tt := range tests {
		res := arraysstrings.IncreasingTriplet(tt.nums)

		assert.Equal(t, tt.expected, res, "test case %d failed", i+1)
	}
}
