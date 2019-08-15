package amazon

func romanToInt(s string) int {
	mp := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	res := 0
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			res += mp[string(s[i])]
		} else {
			if mp[string(s[i])] < mp[string(s[i+1])] {
				res -= mp[string(s[i])]
			} else {
				res += mp[string(s[i])]
			}
		}
	}
	return res
}
