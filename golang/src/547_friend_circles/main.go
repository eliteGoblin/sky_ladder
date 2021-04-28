package _547_friend_circles

func findCircleNum(M [][]int) int {
	if len(M) == 0 {
		return 0
	}
	visited := make([]bool, len(M))
	res := 0
	for i := range M {
		if visited[i] {
			continue
		}
		dfs(M, i, visited)
		res++
	}
	return res
}

func dfs(M [][]int, i int, visited []bool) {
	visited[i] = true
	for j := range M[i] {
		if M[i][j] == 0 || visited[j] {
			continue
		}
		dfs(M, j, visited)
	}
}

// 用DFS求解无向图的连通区域
// 自己的弱项在于经过一次转换后的意义难以分清, i, j, M[i], M[i][j]分别代表什么
// M[i][j]类型int,其实是bool, i代表node的index, j也代表node index;
