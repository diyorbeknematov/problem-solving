package arraysstrings_test

import (
	arraysstrings "leetcode/arrays-strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseVowels(t *testing.T) {
	tests := []struct {
		s        string
		expected string
	}{
		{
			s:        "IceCreAm",
			expected: "AceCreIm",
		},
		{
			s:        "leetcode",
			expected: "leotcede",
		},
	}

	for i, tt := range tests {
		res := arraysstrings.ReverseVowels(tt.s)

		assert.Equal(t, tt.expected, res, "Test case %d failed", i+1)
	}
}
