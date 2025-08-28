package twopointers

func MaxArea(height []int) int {
	mx := 0

	low, high := 0, len(height)-1

	for low < high {
		area := (high - low) * min(height[low], height[high])
		mx = max(mx, area)

		if height[low] < height[high] {
			low++
		} else if height[low] > height[high] {
			high--
		} else {
			low++
			high--
		}
	}
	return mx
}
