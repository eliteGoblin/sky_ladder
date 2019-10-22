package amazon

func minDistance(word1 string, word2 string) int {
	memo := make([][]int, len(word1))
	for i := range memo {
		memo[i] = make([]int, len(word2))
	}
	return minDFS(word1, 0, word2, 0, memo)
}

func minDFS(word1 string, i1 int, word2 string, i2 int, memo [][]int) int {
	if i1 == len(word1) {
		return len(word2) - i2
	}
	if i2 == len(word2) {
		return len(word1) - i1
	}
	if memo[i1][i2] > 0 {
		return memo[i1][i2]
	}
	if word1[i1] == word2[i2] {
		return minDFS(word1, i1+1, word2, i2+1, memo)
	}
	insertCt := minDFS(word1, i1, word2, i2+1, memo)
	replaceCt := minDFS(word1, i1+1, word2, i2+1, memo)
	deleteCt := minDFS(word1, i1+1, word2, i2, memo)
	memo[i1][i2] = min(insertCt, replaceCt, deleteCt) + 1
	return memo[i1][i2]
}
