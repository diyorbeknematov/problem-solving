package hashmapset

import "sort"

func CloseStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	count1 := make(map[rune]int)
	count2 := make(map[rune]int)

	for _, c := range word1 {
		count1[c]++
	}
	for _, c := range word2 {
		count2[c]++
	}

	if len(count1) != len(count2) {
		return false
	}
	for c := range count1 {
		if _, ok := count2[c]; !ok {
			return false
		}
	}

	freq1 := []int{}
	freq2 := []int{}
	for _, v := range count1 {
		freq1 = append(freq1, v)
	}
	for _, v := range count2 {
		freq2 = append(freq2, v)
	}

	sort.Ints(freq1)
	sort.Ints(freq2)

	for i := range freq1 {
		if freq1[i] != freq2[i] {
			return false
		}
	}

	return true
}
