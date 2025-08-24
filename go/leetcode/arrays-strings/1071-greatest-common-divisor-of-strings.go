package arraysstrings

func GcdOfStrings(str1 string, str2 string) string {
	
	gcdLength := gcd(len(str1), len(str2))
	sub := str1[:gcdLength]

	var maxLen = len(str1)
	if len(str1) < len(str2) {
		maxLen = len(str2)
	} 

	for i := 0; i < maxLen; i += gcdLength {
		if i + gcdLength <= len(str1) && (sub != str1[i : i + gcdLength]) {
			return  ""
		}

		if i + gcdLength <= len(str2) && (sub != str2[i : i + gcdLength]) {
			return  ""
		}
	}

	return sub
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a % b
	}

	return a
}