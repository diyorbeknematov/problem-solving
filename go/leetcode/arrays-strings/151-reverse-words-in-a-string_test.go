package arraysstrings_test

import (
	arraysstrings "leetcode/arrays-strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWords(t *testing.T) {
	tests := []struct {
		s        string
		expected string
	}{
		{
			s:        "the sky is blue",
			expected: "blue is sky the",
		},
		{
			s:        "  hello world  ",
			expected: "world hello",
		},
		{
			s:        "a good   example",
			expected: "example good a",
		},
	}

	for i, tt := range tests {
		res := arraysstrings.ReverseWords(tt.s)

		assert.Equal(t, tt.expected, res, "test case %d failed", i+1)
	}
}
