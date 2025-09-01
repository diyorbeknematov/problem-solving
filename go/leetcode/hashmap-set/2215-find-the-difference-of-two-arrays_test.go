package hashmapset_test

import (
	hashmapset "leetcode/hashmap-set"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDifference(t *testing.T) {
	tests := []struct {
		nums1    []int
		nums2    []int
		expected [][]int
	}{
		{
			nums1:    []int{1, 2, 3},
			nums2:    []int{2, 4, 6},
			expected: [][]int{{1, 3}, {4, 6}},
		},
		{
			nums1:    []int{1, 2, 3, 3},
			nums2:    []int{1, 1, 2, 2},
			expected: [][]int{{3}, {}},
		},
	}

	for i, tt := range tests {
		res := hashmapset.FindDifference(tt.nums1, tt.nums2)

		assert.Equal(t, tt.expected, res, "test case %d failed", i+1)
	}
}
