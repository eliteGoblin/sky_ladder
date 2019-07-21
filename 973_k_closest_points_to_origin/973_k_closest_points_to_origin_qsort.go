package amazon

import "sort"

// sort with qsort
type pointsSlice [][]int

func (points pointsSlice) Len() int {
	return len(points)
}

func (points pointsSlice) Less(i, j int) bool {
	distI := distSquare(points[i][0], points[i][1])
	distJ := distSquare(points[j][0], points[j][1])
	return distI < distJ
}

func (points pointsSlice) Swap(i, j int) {
	points[i], points[j] = points[j], points[i]
}

func distSquare(x, y int) int {
	return x*x + y*y
}
func kClosestQSort(points [][]int, K int) [][]int {
	pointsSort := pointsSlice(points)
	sort.Sort(pointsSort)
	return [][]int(pointsSort)[:K]
}
