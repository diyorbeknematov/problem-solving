package hashmapset_test

import (
	hashmapset "leetcode/hashmap-set"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueOccurrences(t *testing.T) {
	tests := []struct {
		arr      []int
		expected bool
	}{
		{
			arr:      []int{1, 2, 2, 1, 1, 3},
			expected: true,
		},
		{
			arr:      []int{1, 2},
			expected: false,
		},
		{
			arr:      []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0},
			expected: true,
		},
	}

	for i, tt := range tests {
		res := hashmapset.UniqueOccurrences(tt.arr)

		assert.Equal(t, tt.expected, res, "test case %d failed", i+1)
	}
}
