package arraysstrings_test

import (
	arraysstrings "leetcode/arrays-strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeAlternately(t *testing.T) {
	word1 := "abcd"
	word2 := "pq"
	expect := "apbqcd"

	res := arraysstrings.MergeAlternately(word1, word2)

	assert.Equal(t, expect, res, "they should be equal")
}

	