package _78_subsets

import "sort"

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	dfs(&res, []int{}, nums)
	return res
}

func dfs(res *[][]int, prefix []int, nums []int) {
	if len(nums) == 0 {
		*res = append(*res, prefix)
		return
	}
	dfs(res, append(append([]int{}, prefix...), nums[0]), nums[1:])
	dfs(res, append([]int{}, prefix...), nums[1:])
}
