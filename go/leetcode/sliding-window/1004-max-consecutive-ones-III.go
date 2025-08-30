package slidingwindow

func LongestOnes(nums []int, k int) int {
	cnt := 0
	left, right := 0, 0

	mx := 0

	for right < len(nums) {
		if nums[right] == 0 {
			cnt++
		}

		for cnt > k {
			if nums[left] == 0 {
				cnt--
			}
			left++
		}

		// bu yerda right-left+1 to‘g‘ri bo‘ladi
		mx = max(mx, right-left+1)
		right++
	}

	return mx
}
