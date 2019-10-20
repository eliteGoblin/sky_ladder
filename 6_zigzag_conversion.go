package amazon

func convert(s string, numRows int) string {
	res := make([]string, numRows)
	for i := 0; i < len(s); {
		for j := 0; i < len(s) && j < len(res); {
			res[j] += s[i : i+1]
			j++
			i++
		}
		for j := len(res) - 2; j > 0 && i < len(s); {
			res[j] += s[i : i+1]
			j--
			i++
		}
	}
	resS := ""
	for _, v := range res {
		resS += v
	}
	return resS
}
