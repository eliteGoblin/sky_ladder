package _64_minimum_path_sum

import "github.com/eliteGoblin/sky_ladder/common"

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	row := len(grid)
	col := len(grid[0])

	sum := make([][]int, row)
	for i := range sum {
		sum[i] = make([]int, col)
		for j := range sum[i] {
			sum[i][j] = -1
		}
	}
	sum[0][0] = grid[0][0]
	return minSum(grid, sum, row-1, col-1)
}

func minSum(grid [][]int, sum [][]int, i, j int) int {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return common.IntMax
	}
	if sum[i][j] >= 0 {
		return sum[i][j]
	}
	sum[i][j] = grid[i][j] + common.Min(
		minSum(grid, sum, i-1, j),
		minSum(grid, sum, i, j-1),
	)
	return sum[i][j]
}
