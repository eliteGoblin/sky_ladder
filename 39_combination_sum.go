package amazon

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	dfsCombinationSum(candidates, 0, []int{}, target, &res)
	return res
}

func dfsCombinationSum(candidates []int, start int, path []int, target int, res *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		*res = append(*res, append([]int{}, path...))
		return
	}
	for i := start; i < len(candidates); i++ {
		path = append(path, candidates[i])
		dfsCombinationSum(candidates, i, path, target-candidates[i], res)
		path = path[:len(path)-1]
	}
}
