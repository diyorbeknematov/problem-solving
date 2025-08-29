package slidingwindow_test

import (
	slidingwindow "leetcode/sliding-window"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxVowels(t *testing.T) {
	tests := []struct {
		s        string
		k        int
		expected int
	}{
		{
			s:        "abciiidef",
			k:        3,
			expected: 3,
		},
		{
			s:        "aeiou",
			k:        2,
			expected: 2,
		},
		{
			s:        "leetcode",
			k:        3,
			expected: 2,
		},
	}

	for i, tt := range tests {
		res := slidingwindow.MaxVowels(tt.s, tt.k)

		assert.Equal(t, tt.expected, res, "test case %d failed", i+1)
	}
}
