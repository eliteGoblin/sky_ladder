package amazon

func minWindow(s string, t string) string {
	if s == "" || t == "" {
		return ""
	}
	mp := make(map[byte]int)
	for i := range t {
		mp[t[i]]++
	}
	matchCount := 0
	left, minLen := 0, intMax
	res := ""
	for i := range s {
		mp[s[i]]--
		if mp[s[i]] > 0 {
			matchCount++
		}
		for matchCount == len(t) {
			if minLen > i-left+1 {
				minLen = i - left + 1
				res = s[left : i+1]
			}
			mp[s[left]]++
			if mp[s[left]] > 0 {
				matchCount--
			}
			left++
		}
	}
	return res
}
