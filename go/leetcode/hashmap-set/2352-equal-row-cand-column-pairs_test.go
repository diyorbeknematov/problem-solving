package hashmapset_test

import (
	hashmapset "leetcode/hashmap-set"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqualPairs(t *testing.T) {
	tests := []struct {
		grid     [][]int
		expected int
	}{
		{
			grid:     [][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}},
			expected: 1,
		},
		{
			grid:     [][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}},
			expected: 3,
		},
	}

	for i, tt := range tests {
		res := hashmapset.EqualPairs(tt.grid)

		assert.Equal(t, tt.expected, res, "test case %d failed", i+1)
	}
}
