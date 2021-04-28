package _417_pacific_atlantic_water_flow

import "github.com/eliteGoblin/sky_ladder/common"

func pacificAtlantic(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return [][]int{}
	}
	m, n := len(matrix), len(matrix[0])
	canReachPacific := createMatrix(m, n)
	canReachAtlantic := createMatrix(m, n)
	for j := 0; j < n; j++ {
		dfs(matrix, 0, j, common.IntMin, canReachPacific)
		dfs(matrix, m-1, j, common.IntMin, canReachAtlantic)
	}
	for i := 0; i < m; i++ {
		dfs(matrix, i, 0, common.IntMin, canReachPacific)
		dfs(matrix, i, n-1, common.IntMin, canReachAtlantic)
	}
	var res [][]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if canReachPacific[i][j] == 1 && canReachAtlantic[i][j] == 1 {
				res = append(res, []int{
					i, j,
				})
			}
		}
	}
	return res
}

var direction = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func dfs(matrix [][]int, i, j int, pre int, canReach [][]int) {
	m, n := len(matrix), len(matrix[0])
	if i < 0 || i >= m || j < 0 || j >= n || canReach[i][j] == 1 || matrix[i][j] < pre {
		return
	}
	canReach[i][j] = 1
	for _, dir := range direction {
		dfs(matrix, i+dir[0], j+dir[1], matrix[i][j], canReach)
	}
}

func createMatrix(m, n int) [][]int {
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}
	return res
}
