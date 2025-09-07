package stack_test

import (
	"leetcode/stack"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAsteroidCollision(t *testing.T) {
	tests := []struct {
		asteroids []int
		expected  []int
	}{
		{
			asteroids: []int{5, 10, -5},
			expected:  []int{5, 10},
		},
		{
			asteroids: []int{8, -8},
			expected:  []int{},
		},
		{
			asteroids: []int{10, 2, -5},
			expected:  []int{10},
		},
		{
			asteroids: []int{2, 3, 4, -5, 1, 3},
			expected:  []int{-5, 1, 3},
		},
		{
			asteroids: []int{-2, -2, 1, -2},
			expected:  []int{-2, -2, -2},
		},
	}

	for i, tt := range tests {
		res := stack.AsteroidCollision(tt.asteroids)

		assert.Equal(t, tt.expected, res, "test case %d failed", i+1)
	}
}
