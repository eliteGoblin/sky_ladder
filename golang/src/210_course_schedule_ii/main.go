package _210_course_schedule_ii

import "container/list"

func findOrder(numCourses int, prerequisites [][]int) []int {
	in := make([]int, numCourses)
	graphAdjList := make([][]int, numCourses)
	for _, v := range prerequisites {
		graphAdjList[v[1]] = append(graphAdjList[v[1]], v[0])
		in[v[0]]++
	}
	q := list.New()
	for i, v := range in {
		if v == 0 {
			q.PushBack(i)
		}
	}
	res := make([]int, 0)
	for q.Len() > 0 {
		head := q.Front().Value.(int)
		res = append(res, head)
		q.Remove(q.Front())

		for _, v := range graphAdjList[head] {
			in[v]--
			if in[v] == 0 {
				q.PushBack(v)
			}
		}
	}
	if len(res) < numCourses {
		return []int{}
	}
	return res
}
