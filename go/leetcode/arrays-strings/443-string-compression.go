package arraysstrings

import "fmt"

func Compress(chars []byte) int {
	k := 0
	cnt := 1

	for i := 1; i <= len(chars); i++ {
		if i < len(chars) && chars[i-1] == chars[i] {
			cnt++
		} else {
			chars[k] = chars[i-1]
			k ++
			if cnt != 1 {
				for _, c := range fmt.Sprint(cnt) {
					chars[k] = byte(c)
					k++
				}
                cnt = 1
			}
		}
	}

	return k
}
