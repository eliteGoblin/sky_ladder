package _38_count_and_say

import "fmt"

func countAndSay(n int) string {
	res := make([]string, 30)
	res[0] = "1"
	for i := 1; i <= 29; i++ {
		next := generate(res[i-1])
		res[i] = next
	}
	return res[n]
}

func generate(s string) string {
	res := ""
	for i := 0; i < len(s); {
		ct := 1
		for j := i + 1; j < len(s); j++ {
			if s[j] != s[i] {
				break
			}
			ct++
		}
		res += fmt.Sprintf("%d%s", ct, string(s[i]))
		i += ct
	}
	return res
}
