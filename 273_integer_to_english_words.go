package amazon

func numberToWords(num int) string {
	res := convertHundred(num % 1000)
	v := []string{"Thousand", "Million", "Billion"}
	for i := 0; i < 3; i++ {
		num /= 1000
		if num%1000 > 0 {
			res = convertHundred(num%1000) + " " + v[i]
		}
	}
	if res == "" {
		return "Zero"
	}
	return res
}

func convertHundred(num int) string {
	under20 := []string{
		"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen",
	}
	decimal := []string{
		"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety",
	}
	res := ""
	a, b, c := num/100, num%100, num%10
	if b < 20 {
		res = under20[b]
	} else {
		res = decimal[b/10]
		if c > 0 {
			res += " " + under20[c]
		}
	}
	if a > 0 {
		res = decimal[a] + " Hundred " + res
	}
	return res
}
