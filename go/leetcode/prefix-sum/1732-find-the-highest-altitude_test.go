package prefixsum_test

import (
	prefixsum "leetcode/prefix-sum"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestAltitude(t *testing.T) {
	tests := []struct {
		gain     []int
		expected int
	}{
		{
			gain:     []int{-5, 1, 5, 0, -7},
			expected: 1,
		},
		{
			gain:     []int{-4, -3, -2, -1, 4, 3, 2},
			expected: 0,
		},
	}

	for i, tt := range tests {
		res := prefixsum.LargestAltitude(tt.gain)

		assert.Equal(t, tt.expected, res, "test case %d failed", i+1)
	}
}
