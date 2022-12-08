package _93_restore_ip_addresses

import "strings"

func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	helper(s, []string{}, &res)
	return res
}

func helper(s string, path []string, res *[]string) {
	if s == "" && len(path) == 4 {
		*res = append(*res, strings.Join(path, "."))
		return
	}
	if s == "" {
		return
	}
	if len(path) >= 4 {
		return
	}
	helper(s[1:], append(path, s[:1]), res)
	if len(s) == 1 {
		return
	}
	if s[0] == '0' {
		return
	}
	helper(s[2:], append(path, s[:2]), res)
	if len(s) == 2 || s[:3] > "255" {
		return
	}
	helper(s[3:], append(path, s[:3]), res)
}
