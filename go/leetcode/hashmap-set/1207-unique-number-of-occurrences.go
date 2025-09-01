package hashmapset

func UniqueOccurrences(arr []int) bool {
	freq := make(map[int]int)
	for _, num := range arr {
		freq[num]++
	}

	isTrue := make(map[int]bool)

	for _, val := range freq {
		if isTrue[val] {
			return false
		}

		isTrue[val] = true
	}

	return true
}
