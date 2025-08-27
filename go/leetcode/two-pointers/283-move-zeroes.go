package twopointers

func MoveZeroes(nums []int) {
	n := len(nums)
	i, j := 0, 0

	for j < n {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i ++
		}

		j ++
	}
}

// 0, 1, 0, 3, 12 | i = 0, j = 0
// 0, 1, 0, 3, 12 | i = 0, j = 1
// 1, 0, 0, 3, 12 | i = 1, j = 2
// 1, 0, 0, 3, 12 | i = 1, j = 3
// 1, 3, 0, 0, 12 | i = 2, j = 4
// 1, 3, 12, 0, 0 | i = 3, j = 5