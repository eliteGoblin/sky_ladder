package _207_course_schedule

const (
	nonvisited = iota
	isvisited
	conflict
)

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	for _, v := range prerequisites {
		graph[v[1]] = append(graph[v[1]], v[0])
	}
	visited := make([]int, numCourses)
	for i := 0; i < numCourses; i++ {
		if isCycle(graph, i, visited) {
			return false
		}
	}
	return true
}

func isCycle(graph [][]int, i int, visited []int) bool {
	if visited[i] == conflict {
		return true
	}
	if visited[i] == isvisited {
		return false
	}
	visited[i] = conflict
	for _, v := range graph[i] {
		if isCycle(graph, v, visited) {
			return true
		}
	}
	visited[i] = isvisited
	return false
}

// 判断Graph是否有环: 遍历每个节点: 有bfs(利用出度数组)和dfs:利用visited, dfs过程中发现visit过的，即为环
// 巧妙之处:  不必每次遍历节点都建立visited, 而且可以缓存之前判定为无环的节点结果
// visited 3状态: 0没有访问, 1访问过，无环 2访问过，有环;
