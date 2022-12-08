package _387_first_unique_character_in_a_string

func firstUniqChar(s string) int {
	mp := make(map[byte]int)
	for i := range s {
		mp[s[i]]++
	}
	for i := range s {
		if mp[s[i]] == 1 {
			return i
		}
	}
	return -1
}
