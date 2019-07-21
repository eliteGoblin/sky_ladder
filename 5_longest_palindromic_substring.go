package amazon

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	var start, end int
	for i := 0; i < len(s); i++ {
		curS, curE := findLongestPalindrome(s, i, i)
		if curE-curS > end-start {
			start, end = curS, curE
		}
		curS, curE = findLongestPalindrome(s, i, i+1)
		if curE-curS > end-start {
			start, end = curS, curE
		}
	}
	return s[start:end]
}

func findLongestPalindrome(s string, left, right int) (start, end int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right
}
