package slidingwindow

func MaxVowels(s string, k int) int {
	isVowel := func(ch byte) bool {
		return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
	}

	cnt := 0
	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			cnt++
		}
	}

	mx := cnt
	for i := k; i < len(s); i++ {
		if isVowel(s[i-k]) {
			cnt--
		}
		if isVowel(s[i]) {
			cnt++
		}
		if cnt > mx {
			mx = cnt
		}
	}

	return mx
}
