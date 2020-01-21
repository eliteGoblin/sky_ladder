package amazon

import "fmt"

func exist(board [][]byte, word string) bool {
	for i := range board {
		for j := range board[i] {
			if contains(word, board, i, j) {
				return true
			}
		}
	}
	return false
}

var move = [][2]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func contains(word string, board [][]byte, i, j int) bool {
	if word == "" {
		fmt.Println("found")
		return true
	}
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return false
	}
	if board[i][j] == '#' || word[0] != board[i][j] {
		return false
	}
	board[i][j] = '#'
	res := false
	for k := range move {
		res = contains(word[1:], board, i+move[k][0], j+move[k][1])
		if res {
			break
		}
	}
	board[i][j] = word[0]
	return res
}
