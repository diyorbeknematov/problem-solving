package prefixsum

func PivotIndex(nums []int) int {
	leftSum := 0
	totalSum := 0

	for _, num := range nums {
		totalSum += num
	}

	for i := 0; i < len(nums); i++ {
		rightSum := totalSum - leftSum - nums[i]
		if rightSum == leftSum {
			return i
		}

		leftSum += nums[i]
	}

	return -1
}
