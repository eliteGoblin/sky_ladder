package _529_minesweeper

var dir = [][2]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func updateBoard(board [][]byte, click []int) [][]byte {
	x, y := click[0], click[1]
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) {
		return board
	}
	if board[x][y] == 'M' {
		board[x][y] = 'X'
		return board
	}
	if board[x][y] != 'E' {
		return board
	}
	adj := searchAdj(board, x, y)
	if adj == 0 {
		board[x][y] = 'B'
		for i := range dir {
			r, c := x+dir[i][0], y+dir[i][1]
			updateBoard(board, []int{r, c})
		}
	} else {
		board[x][y] = '0' + adj
	}
	return board
}

func searchAdj(board [][]byte, x, y int) byte {
	var ct byte
	for i := range dir {
		r, c := x+dir[i][0], y+dir[i][1]
		if r < 0 || r >= len(board) || c < 0 || c >= len(board[0]) {
			continue
		}
		if board[r][c] == 'M' {
			ct++
		}
	}
	return ct
}
