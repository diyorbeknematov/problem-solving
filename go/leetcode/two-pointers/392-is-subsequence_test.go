package twopointers_test

import (
	twopointers "leetcode/two-pointers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		s, t     string
		expected bool
	}{
		{
			s:        "abc",
			t:        "ahbgdc",
			expected: true,
		},
		{
			s:        "axc",
			t:        "ahbgdc",
			expected: false,
		},
	}

	for i, tt := range tests {
		res := twopointers.IsSubsequence(tt.s, tt.t)

		assert.Equal(t, tt.expected, res, "test case %d failed", i+1)
	}
}
