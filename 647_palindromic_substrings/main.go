package _647_palindromic_substrings

func countSubstrings(s string) int {
	if len(s) == 0 {
		return 0
	}
	var res int
	for i := 0; i < len(s); i++ {
		res += helper(s, i, i)
		res += helper(s, i, i+1)
	}
	return res
}

func helper(s string, i, j int) int {
	var res int
	for i >= 0 && j < len(s) && s[i] == s[j] {
		res++
		i++
		j--
	}
	return res
}
