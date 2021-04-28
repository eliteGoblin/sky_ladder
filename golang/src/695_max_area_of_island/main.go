package _695_max_area_of_island

func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	for i := range grid {
		for j := range grid[i] {
			area := islandArea(grid, i, j)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func islandArea(grid [][]int, i, j int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return 0
	}
	if grid[i][j] != 1 {
		return 0
	}
	grid[i][j] = -1
	res := 1
	res += islandArea(grid, i-1, j)
	res += islandArea(grid, i, j-1)
	res += islandArea(grid, i+1, j)
	res += islandArea(grid, i, j+1)
	return res
}
