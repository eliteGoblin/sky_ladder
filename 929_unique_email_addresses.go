package amazon

import "strings"

func numUniqueEmails(emails []string) int {
	mp := make(map[string]int)
	for _, v := range emails {
		addrs := strings.Split(v, "@")
        res := ""
		FOR:
		for i := range addrs[0] {
			switch addrs[0][i] {
			case '.':
				continue
			case '+':
				break FOR
			default:
				res += string(addrs[0][i])
			}
		}
		res += "@" + addrs[1]
		mp[res] ++
	}
	return len(mp)
}