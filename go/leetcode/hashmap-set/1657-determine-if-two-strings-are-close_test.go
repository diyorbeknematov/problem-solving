package hashmapset_test

import (
	hashmapset "leetcode/hashmap-set"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCloseStrings(t *testing.T) {
	tests := []struct {
		word1    string
		word2    string
		expected bool
	}{
		{
			word1:    "abc",
			word2:    "bca",
			expected: true,
		},
		{
			word1:    "a",
			word2:    "aa",
			expected: false,
		},
		{
			word1:    "cabbba",
			word2:    "abbccc",
			expected: true,
		},
	}

	for i, tt := range tests {
		res := hashmapset.CloseStrings(tt.word1, tt.word2)

		assert.Equal(t, tt.expected, res, "test case %d failed", i+1)
	}
}
