package arraysstrings

import (
	"slices"
	"strings"
)

func ReverseWords(s string) string {
    // var (
	// 	ans string 
	// 	word string
	// )

	// s = strings.TrimSpace(s)

	// for _, c := range s {
	// 	if c == ' ' && word != "" {
	// 		ans = word + " " + ans
	// 		word = ""
	// 	} else if c != ' ' {
	// 		word += string(c)
	// 	}
	// }

	// ans = strings.TrimSpace(string(word) + " " + ans)

	// return ans

	// method 2
	words := strings.Fields(s)
	slices.Reverse(words)

	return strings.Join(words, " ")

}