package amazon

import (
	"sort"
)

type intervals [][]int

func (a intervals) Less(i, j int) bool {
	return a[i][0] < a[j][0]
}

func (a intervals) Len() int {
	return len(a)
}

func (a intervals) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func merge(ins [][]int) [][]int {
	if len(ins) <= 1 {
		return ins
	}
	intervalArr := intervals(ins)
	sort.Sort(intervalArr)
	res := make([][]int, 0)
	res = append(res, intervalArr[0])
	for i := 1; i < len(intervalArr); i++ {
		if intervalArr[i][0] <= res[len(res)-1][1] {
			res[len(res)-1][1] = max(res[len(res)-1][1], intervalArr[i][1])
		} else {
			res = append(res, intervalArr[i])
		}
	}
	return res
}
