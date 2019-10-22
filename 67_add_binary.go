package amazon

func addBinary(a string, b string) string {
	res := ""
	i, j := len(a)-1, len(b)-1
	var carry byte
	for i >= 0 || j >= 0 {
		var va, vb, sum byte
		if i >= 0 {
			va = a[i] - '0'
			i--
		}
		if j >= 0 {
			vb = b[j] - '0'
			j--
		}
		sum = va + vb + carry
		res = string(sum%2+'0') + res
		carry = sum / 2
	}
	if carry == 1 {
		res = "1" + res
	}
	return res
}

func addBinarySelf(a string, b string) string {
	res := ""
	i, j := len(a)-1, len(b)-1
	var carry byte = 0
	for i >= 0 && j >= 0 {
		v := a[i] + b[j] - '0' - '0' + carry
		carry = v / 2
		res += string(byte(v%2 + '0'))

		i--
		j--
	}
	for i >= 0 {
		v := a[i] - '0' + carry
		carry = v / 2
		res += string(byte(v%2 + '0'))
		i--
	}
	for j >= 0 {
		v := a[j] - '0' + carry
		carry = v / 2
		res += string(byte(v%2 + '0'))
		j--
	}
	r := []rune(res)
	i, j = 0, len(r)-1
	for i < j {
		r[i], r[j] = r[j], r[i]
		i++
		j--
	}
	return string(r)
}
