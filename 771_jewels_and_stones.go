package amazon

func numJewelsInStones(J string, S string) int {
	mp := make(map[byte]bool)
	for i := range J {
		mp[J[i]] = true
	}
	res := 0
	for i := range S {
		if _, ok := mp[S[i]]; ok {
			res++
		}
	}
	return res
}
