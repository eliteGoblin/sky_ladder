package amazon

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	revert := 0
	org := x
	for org > 0 {
		revert = revert*10 + org%10
		org = org / 10
	}
	return revert == x
}
