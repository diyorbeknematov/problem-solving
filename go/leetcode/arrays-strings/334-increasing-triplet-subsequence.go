package arraysstrings

import "math"

func IncreasingTriplet(nums []int) bool {
    first, second := math.MaxInt64, math.MaxInt64

	for _, num := range nums {
		if first >= num {
			first = num
		} else if second >= num {
			second = num
		} else {
			return true
		}
	}
	return false
}