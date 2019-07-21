package amazon

func lengthOfLongestSubstring(s string) int {
	sB := []byte(s)
	mp := make(map[byte]int)
	prePos := -1
	var res int
	for i := range s {
		if pos, ok := mp[sB[i]]; ok {
			if pos > prePos {
				// duplicate found
				mp[sB[i]] = i
				prePos = pos + 1
			}
		}
		if i-prePos > res {
			res = i - prePos
		}
	}
	return res
}
