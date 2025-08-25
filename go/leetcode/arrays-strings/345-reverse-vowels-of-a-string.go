package arraysstrings

func ReverseVowels(s string) string {
	vowels := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}

	b := []byte(s)
	i, j := 0, len(s)-1

	for i < j {
		for i < j && !vowels[b[i]] {
			i++
		}
		for i < j && !vowels[b[j]] {
			j--
		}

		b[i], b[j] = b[j], b[i]
		i++
		j--
	}

	return string(b)
}
