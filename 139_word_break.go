package amazon

func wordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]int)
	memo := make(map[string]bool)
	for _, v := range wordDict {
		wordSet[v] = 1
	}
	return wordBreakDFS(s, wordSet, memo)
}

func wordBreakDFS(s string, wordSet map[string]int, memo map[string]bool) bool {
	if s == "" {
		return true
	}
	if v, ok := memo[s]; ok {
		return v
	}
	for i := 1; i <= len(s); i++ {
		prefix := s[:i]
		if _, ok := wordSet[prefix]; ok {
			if wordBreakDFS(s[i:], wordSet, memo) {
				memo[s] = true
				return memo[s]
			}
		}
	}
	memo[s] = false
	return memo[s]
}
