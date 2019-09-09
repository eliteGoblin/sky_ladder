package amazon

func decodeString(s string) string {
	var pos int
	return decode(s, &pos)
}

func decode(s string, pos *int) string {
	res := ""
	for *pos < len(s) && s[*pos] != ']'
}
