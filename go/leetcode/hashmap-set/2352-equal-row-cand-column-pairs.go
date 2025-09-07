package hashmapset

import (
	"strconv"
	"strings"
)

func EqualPairs(grid [][]int) int {
	n := len(grid)

	rowCount := make(map[string]int)
	for i := 0; i < n; i++ {
		parts := make([]string, n)
		for j := 0; j < n; j++ {
			parts[j] = strconv.Itoa(grid[i][j])
		}
		key := strings.Join(parts, ",")
		rowCount[key]++
	}

	count := 0
	for j := 0; j < n; j++ {
		parts := make([]string, n)
		for i := 0; i < n; i++ {
			parts[i] = strconv.Itoa(grid[i][j])
		}
		key := strings.Join(parts, ",")
		if val, ok := rowCount[key]; ok {
			count += val
		}
	}

	return count
}
