package arraysstrings

func MergeAlternately(word1 string, word2 string) string {
    ans := make([]byte, 0, len(word1) + len(word2))

	for i, j := 0, 0; i < len(word1) || j < len(word2); {
		if i < len(word1) {
			ans = append(ans, word1[i])
			i++
		}
		if j < len(word2) {
			ans = append(ans, word2[j])
			j++
		}
	}

	return string(ans)
}