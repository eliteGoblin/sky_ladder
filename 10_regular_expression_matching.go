package amazon

func isMatch(s string, p string) bool {
	if p == "" {
		return s == ""
	}
	if len(p) == 1 {
		return len(s) == 1 && matchChar(s[0], p[0])
	}
	if p[1] == '*' {
		return isMatch(s, p[2:]) || (s != "" && matchChar(s[0], p[0]) && isMatch(s[1:], p))
	}
	return s != "" && matchChar(s[0], p[0]) && isMatch(s[1:], p[1:])
}

func matchChar(sc, pc byte) bool {
	if pc == '.' {
		return true
	}
	return sc == pc
}
