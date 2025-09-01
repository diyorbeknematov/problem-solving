package hashmapset

func FindDifference(nums1 []int, nums2 []int) [][]int {
	mp := make(map[int]int)

	for _, num := range nums1 {
		mp[num] = 1
	}

	for _, num := range nums2 {
		if mp[num] == 1 {
			mp[num] = 2
		} else if mp[num] == 0 {
			mp[num] = -1
		}
	}

	ans := [][]int{make([]int, 0), make([]int, 0)}

	for key, val := range mp {
		if val == 1 {
			ans[0] = append(ans[0], key)
		} else if val == -1 {
			ans[1] = append(ans[1], key)
		}
	}

	return ans
}
