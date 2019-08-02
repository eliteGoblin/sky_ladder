package amazon

func gameOfLife(board [][]int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			ct := getLiveNeighbours(board, i, j)
			if board[i][j] == 0 {
				if ct == 3 {
					board[i][j] = 3
				}
			} else {
				if ct < 2 || ct > 3 {
					board[i][j] = 2
				}
			}
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			board[i][j] %= 2
		}
	}
}

var deltaX = []int{-1, -1, -1, 0, 0, 1, 1, 1}
var deltaY = []int{-1, 0, 1, -1, 1, -1, 0, 1}

func getLiveNeighbours(board [][]int, x, y int) int {
	ct := 0
	for k := 0; k < 8; k++ {
		i := x + deltaX[k]
		j := y + deltaY[k]
		if i >= 0 && i < len(board) && j >= 0 && j < len(board[0]) {
			if board[i][j] == 1 || board[i][j] == 2 {
				ct++
			}
		}
	}
	return ct
}
