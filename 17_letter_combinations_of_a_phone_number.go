package amazon

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	mp := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	res := make([]string, 0)
	letterDFS(digits, "", mp, &res)
	return res
}

func letterDFS(digits string, cur string, mp map[byte]string, res *[]string) {
	if len(digits) == 0 {
		*res = append(*res, cur)
		return
	}
	s := mp[digits[0]]
	for i := range s {
		letterDFS(digits[1:], cur+string(s[i]), mp, res)
	}
}
