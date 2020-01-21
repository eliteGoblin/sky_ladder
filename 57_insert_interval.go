package amazon

func insert(intervals [][]int, newInterval []int) [][]int {
	cur := 0
	res := make([][]int, 0)
	for cur < len(intervals) &&
		(intervals[cur][1] < newInterval[0]) {
		res = append(res, intervals[cur])
		cur++
	}
	for cur < len(intervals) && newInterval[1] >= intervals[cur][0] {
		newInterval[0] = min(newInterval[0], intervals[cur][0])
		newInterval[1] = max(newInterval[1], intervals[cur][1])
		cur++
	}
	res = append(res, newInterval)
	for cur < len(intervals) {
		res = append(res, intervals[cur])
		cur++
	}
	return res
}

// [[1,3],[6,9]]
// [2,5]
