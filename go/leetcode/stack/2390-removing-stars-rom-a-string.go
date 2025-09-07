package stack

func RemoveStars(s string) string {
	var stk []rune

	for _, c := range s {
		if c == '*' {
			stk = stk[:len(stk)-1]
		} else {
			stk = append(stk, c)
		}
	}

	return string(stk)
}
