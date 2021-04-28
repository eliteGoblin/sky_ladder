package _25_valid_palindrome

import "strings"

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	i, j := 0, len(s) - 1
	for i < j {
		if !isAlphaNum(s[i]){
			i ++
			continue
		}
		if !isAlphaNum(s[j]){
			j --
			continue
		}
		if s[i] != s[j] {
			return false
		}
		i ++
		j --
	}
	return true
}

func isAlphaNum(b byte)bool{
	if b >= 'a' && b <= 'z' {
		return true
	}
	if b >= '0' && b <= '9' {
		return true
	}
	return false
}