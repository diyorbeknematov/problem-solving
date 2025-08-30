package slidingwindow

func LongestSubarray(nums []int) int {
	cnt := 0
	mx := 0

	left, right := 0, 0
	for right < len(nums) {
		if nums[right] == 0 {
			cnt++
		}

		if cnt > 1 {
			if nums[left] == 0 {
				cnt--
			}
			left++
		}

		mx = max(mx, right-left)
		right++
	}

	return mx
}
