package amazon

func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				count++
				markIslandContainLocation(grid, i, j)
			}
		}
	}
	return count
}

func markIslandContainLocation(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) ||
		grid[i][j] != '1' {
		return
	}
	grid[i][j] = '2'
	markIslandContainLocation(grid, i-1, j)
	markIslandContainLocation(grid, i+1, j)
	markIslandContainLocation(grid, i, j-1)
	markIslandContainLocation(grid, i, j+1)
}
