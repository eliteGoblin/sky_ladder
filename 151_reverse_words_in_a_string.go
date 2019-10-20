package amazon

import "strings"

func reverseWords(s string) string {
	arr := strings.Split(s, " ")
	nonEmpty := make([]string, 0, len(arr))
	for _, v := range arr {
		if v != "" {
			nonEmpty = append(nonEmpty, v)
		}
	}
	if len(nonEmpty) == 0 {
		return ""
	}
	res := ""
	for i := len(nonEmpty); i >= 0; i-- {
		res += nonEmpty[i] + " "
	}
	return res[:len(res)-1]
}
