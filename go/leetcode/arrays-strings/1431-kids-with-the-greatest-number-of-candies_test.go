package arraysstrings_test

import (
	arraysstrings "leetcode/arrays-strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKidsWithCandies(t *testing.T) {
	tests := []struct {
		candies      []int
		extraCandies int
		expected     []bool
	}{
		{
			candies:      []int{2, 3, 5, 1, 3},
			extraCandies: 3,
			expected:     []bool{true, true, true, false, true},
		},
		{
			candies:      []int{4, 2, 1, 1, 2},
			extraCandies: 1,
			expected:     []bool{true, false, false, false, false},
		},
	}

	for i, tt := range tests {
		res := arraysstrings.KidsWithCandies(tt.candies, tt.extraCandies)
		assert.Equal(t, tt.expected, res, "Test case %d failed", i+1)
	}
}
