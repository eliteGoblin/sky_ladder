package amazon

func canFinish(numCourses int, prerequisites [][]int) bool {
	if numCourses <= 1 {
		return true
	}
	in := make([]int, numCourses)
	graph := make([][]int, numCourses)
	for i := range graph {
		graph[i] = make([]int, 0, numCourses)
	}
	for _, v := range prerequisites {
		in[v[0]]++
		graph[v[1]] = append(graph[v[1]], v[0])
	}
	q := make([]int, 0)
	for i, v := range in {
		if v == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		e := q[0]
		q = q[1:]
		for _, v := range graph[e] {
			in[v]--
			if in[v] == 0 {
				q = append(q, v)
			}
		}
	}
	for _, v := range in {
		if v > 0 {
			return false
		}
	}
	return true
}
