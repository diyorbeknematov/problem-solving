package arraysstrings_test

import (
	arraysstrings "leetcode/arrays-strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompress(t *testing.T) {
	tests := []struct {
		chars    []byte
		expected int
	}{
		{
			chars:    []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
			expected: 6,
		},
		{
			chars:    []byte{'a'},
			expected: 1,
		},
		{
			chars:    []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
			expected: 4,
		},
	}

	for i, tt := range tests {
		res := arraysstrings.Compress(tt.chars)

		assert.Equal(t, tt.expected, res, "test case %d failed", i+1)
	}
}
