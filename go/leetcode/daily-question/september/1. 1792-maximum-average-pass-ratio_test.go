package september_test

import (
	"leetcode/daily-question/september"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxAverageRatio(t *testing.T) {
	tests := []struct {
		classes       [][]int
		extraStudents int
		expected      float64
	}{
		{
			classes:       [][]int{{1, 2}, {3, 5}, {2, 2}},
			extraStudents: 2,
			expected:      0.78333,
		},
		{
			classes:       [][]int{{2, 4}, {3, 9}, {4, 5}, {2, 10}},
			extraStudents: 4,
			expected:      0.53485,
		},
	}

	for i, tt := range tests {
		res := september.MaxAverageRatio(tt.classes, tt.extraStudents)

		assert.Equal(t, tt.expected, res, "test case %d failed", i+1)
	}
}
