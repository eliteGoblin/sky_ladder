package amazon

import "fmt"

func groupAnagrams(strs []string) [][]string {
	mp := make(map[string][]string)
	for _, v := range strs {
		key := countChar(v)
		mp[key] = append(mp[key], v)
	}
	res := make([][]string, 0)
	for k := range mp {
		res = append(res, mp[k])
	}
	return res
}

func countChar(s string) string {
	btMap := make([]byte, 26)
	for i := range s {
		btMap[s[i]-'a']++
	}
	res := ""
	for i := range btMap {
		res += fmt.Sprintf("%d ", btMap[i])
	}
	return res
}
