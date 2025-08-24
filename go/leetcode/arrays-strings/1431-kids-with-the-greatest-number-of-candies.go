package arraysstrings

func KidsWithCandies(candies []int, extraCandies int) []bool {
    var ans = make([]bool, len(candies))

	mx := candies[0]
	for i := 0; i < len(candies); i ++ {
		mx = max(mx, candies[i])
	}

	for i := 0; i < len(candies); i ++ {
		ans[i] = mx <= candies[i] + extraCandies
	}

	return ans
}