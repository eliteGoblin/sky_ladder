package amazon

func numDecodings(s string) int {
	dict := make(map[string]int)
	dict[""] = 1
	return dfsNumDecoding(s, dict)
}

func dfsNumDecoding(s string, dict map[string]int) int {
	if ct, ok := dict[s]; ok {
		return ct
	}
	res := 0
	if s[0] == '0' {
		return 0
	}
	if len(s) == 1 {
		res = 1
	} else {
		res = dfsNumDecoding(s[1:], dict)
		if s[0] == '1' ||
			(s[0] == '2' && s[1] >= '0' && s[1] <= '6') {
			res += dfsNumDecoding(s[2:], dict)
		}
	}
	dict[s] = res
	return dict[s]
}
