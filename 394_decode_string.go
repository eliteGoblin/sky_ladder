package amazon

func decodeString(s string) string {
	var pos int
	return decode(s, &pos)
}

func decode(s string, i *int) string {
	res := ""
	for *i < len(s) && s[*i] != ']' {
		if s[*i] < '0' || s[*i] > '9' {
			res += string(s[*i])
			*i ++
		} else {
			rep := 0
			for s[*i] >= '0' && s[*i] <= '9' {
				rep = rep * 10 + int(s[*i] - '0')
				*i ++
			}
			*i ++
			sub := decode(s, i)
			*i ++
			for rep > 0 {
				res += sub
				rep --
			}
		}
	}
	return res
}
