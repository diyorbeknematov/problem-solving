package slidingwindow

import "math"

func FindMaxAverage(nums []int, k int) float64 {
	var s int
	for i := 0; i < k; i++ {
		s += nums[i]
	}

	mx := float64(s) / float64(k)

	for i := k; i < len(nums); i++ {
		s = s - nums[i-k] + nums[i]
		avg := float64(s) / float64(k)
		if mx < avg {
			mx = avg
		}
	}

	return round(mx, 5)
}

func round(x float64, prec int) float64 {
	pow := math.Pow(10, float64(prec))
	return math.Round(x*pow) / pow
}

// 1,12,-5,-6,50,3, k = 4
// s = 1 + 12 - 5 - 6 = 2, avg = 2 / 4, mx = max(0, 0.50000)
// s = 1 + 12 - 5 - 6 - 1 + 50, avg = 51 / 4 = 12,75000, mx = max(0.50000, 12.75000)

// 1456-maximum-number-of-vowels-in-substring-length.go
