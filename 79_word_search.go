package amazon

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	for i := range board {
		for j := range board[0] {
			if findDFS(board, i, j, word) {
				return true
			}
		}
	}
	return false
}

func findDFS(board [][]byte, i, j int, word string) bool {
	if word == "" {
		return true
	}
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) ||
		board[i][j] == '#' {
		return false
	}
	if board[i][j] != word[0] {
		return false
	}
	board[i][j] = '#'
	defer func() {
		board[i][j] = word[0]
	}()
	find := findDFS(board, i-1, j, word[1:])
	if find {
		return true
	}
	find = findDFS(board, i+1, j, word[1:])
	if find {
		return true
	}
	find = findDFS(board, i, j-1, word[1:])
	if find {
		return true
	}
	find = findDFS(board, i, j+1, word[1:])
	if find {
		return true
	}
	return false
}
