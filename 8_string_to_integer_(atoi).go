package amazon

const (
	INT32MAX = int(^uint32(0) >> 1)
	INT32MIN = -INT32MAX - 1
)

func myAtoi(str string) int {
	var res int
	var i int
	for i = 0; i < len(str); i++ {
		if str[i] != ' ' {
			break
		}
	}
	if i == len(str) {
		return 0
	}
	sign := 1
	if str[i] == '+' || str[i] == '-' {
		if str[i] == '-' {
			sign = -1
		}
		i++
	}
	for ; i < len(str); i++ {
		if str[i] < '0' || str[i] > '9' {
			break
		}
		if res > INT32MAX/10 || 
        (res == INT32MAX / 10 && str[i] > '7') {
				if sign == 1{
                    return INT32MAX
				}
            return int(INT32MIN)
		}
        res = res * 10 + int(str[i] - '0')
	}
    return res * sign
}
