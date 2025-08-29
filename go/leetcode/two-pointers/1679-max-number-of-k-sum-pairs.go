package twopointers

// import "sort"

func MaxOperations(nums []int, k int) int {
	// sort.Ints(nums)

	// first, end := 0, len(nums)-1
	// ans := 0

	// for first < end {
	// 	if nums[first]+nums[end] == k {
	// 		ans++
	// 		first++
	// 		end--
	// 	} else if nums[first]+nums[end] > k {
	// 		end--
	// 	} else {
	// 		first++
	// 	}
	// }

	freq := make(map[int]int)
	ans := 0

	for _, num := range nums {
		need := k - num
		if freq[need] > 0 {
			ans++
			freq[need]--
		} else {
			freq[num]++
		}
	}

	return ans
}
