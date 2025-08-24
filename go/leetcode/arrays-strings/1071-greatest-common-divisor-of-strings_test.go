package arraysstrings_test

import (
	arraysstrings "leetcode/arrays-strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGcdOfStrings(t *testing.T) {
	tests := []struct {
		str1     string
		str2     string
		expected string
	}{
		{
			str1:     "ABCABC",
			str2:     "ABC",
			expected: "ABC",
		},
		{
			str1:     "LEET",
			str2:     "CODE",
			expected: "",
		},
	}

	for i, tt := range tests {
		res := arraysstrings.GcdOfStrings(tt.str1, tt.str2)

		assert.Equal(t, tt.expected, res, "Test case %d failed", i+1)
	}
}
