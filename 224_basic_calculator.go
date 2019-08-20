package amazon

func calculateRecursive(s string) int {
	num := 0
	res := 0
	sign := 1
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num = num*10 + int(s[i]-'0')
		}
		if s[i] == '(' {
			ct := 0
			start := i
			for ; i < len(s); i++ {
				if s[i] == '(' {
					ct++
				} else if s[i] == ')' {
					ct--
				}
				if ct == 0 {
					break
				}
			}
			num = calculateRecursive(s[start+1 : i])
		}
		if s[i] == '+' || s[i] == '-' || i == len(s)-1 {
			res += sign * num
			num = 0
			if s[i] == '+' {
				sign = 1
			} else if s[i] == '-' {
				sign = -1
			}
		}
	}
	return res
}
