package amazon

import "sort"

func permuteByNext(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for {
		res = append(res, append([]int{}, nums...))
		if !nextPerm(nums) {
			break
		}
	}
	return res
}

func nextPerm(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}
	p := -1
	for i := len(nums) - 1; i >= 1; i-- {
		if nums[i-1] < nums[i] {
			p = i - 1
			break
		}
	}
	if p == -1 {
		return false
	}
	q := -1
	for i := len(nums) - 1; i > p; i-- {
		if nums[i] > nums[p] {
			q = i
			break
		}
	}
	nums[p], nums[q] = nums[q], nums[p]
	sort.Ints(nums[p+1:])
	return true
}

func permute(p []int) [][]int {
	sort.Ints(p)
	visited := make(map[int]bool)
	res := make([][]int, 0)
	permuteDFS(&res, visited, p, []int{})
	return res
}

func permuteDFS(res *[][]int, visited map[int]bool, p []int, cur []int) {
	if len(cur) == len(p) {
		*res = append(*res, append([]int{}, cur...))
		return
	}
	for _, v := range p {
		// if _, ok := visited[v]; ok {  易出错的地方，用bool标记是否存在，不需要delete; 而用ok判断是否存在，则需要delete，容易弄错；
		// 	continue
		// }
		if visited[v] {
			continue
		}
		visited[v] = true
		permuteDFS(res, visited, p, append(cur, v))
		visited[v] = false
	}
}
