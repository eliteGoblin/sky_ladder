package _438_find_all_anagrams_in_a_string

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	mpS := buildSum(s[:len(p)])
	mpP := buildSum(p)
	diff := diff(mpS, mpP)
	var res []int
	if isSame(diff) {
		res = append(res, 0)
	}

	// 这里关键, 下标计算对自己是最难的
	for i := 1; i < len(s)-len(p)+1; i++ {
		diff[s[i-1]-'a']--
		diff[s[i+len(p)-1]-'a']++
		if isSame(diff) {
			res = append(res, i)
		}
	}

	return res
}

func buildSum(s string) [26]byte {
	var mp [26]byte
	for i := range s {
		mp[s[i]-'a']++
	}
	return mp
}

func diff(m1, m2 [26]byte) [26]byte {
	var res [26]byte
	for i := 0; i < 26; i++ {
		res[i] = m1[i] - m2[i]
	}
	return res
}

func isSame(diff [26]byte) bool {
	for i := range diff {
		if diff[i] != 0 {
			return false
		}
	}
	return true
}
