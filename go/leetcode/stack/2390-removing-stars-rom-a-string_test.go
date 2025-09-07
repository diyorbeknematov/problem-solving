package stack_test

import (
	"leetcode/stack"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveStars(t *testing.T) {
	tests := []struct {
		s        string
		expected string
	}{
		{
			s:        "leet**cod*e",
			expected: "lecoe",
		},
		{
			s:        "erase*****",
			expected: "",
		},
	}

	for i, tt := range tests {
		res := stack.RemoveStars(tt.s)

		assert.Equal(t, tt.expected, res, "test case %d failed", i+1)
	}
}
