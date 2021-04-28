package _1_n_queens

func solveNQueens(n int) [][]string {
	queenCol := make([]int, n)
	res := make([][]string, 0)
	helper(0, queenCol, &res)
	return res
}

func helper(curRow int, queenCol []int, res *[][]string){
	if curRow == len(queenCol) {
		*res = append(*res, boardFromQueenCol(queenCol))
		return
	}
	for i := 0; i < len(queenCol); i ++ {
		if isValid(queenCol, curRow, i) {
			queenCol[curRow] = i
			helper(curRow+1, queenCol, res)
			queenCol[curRow] = -1
		}
	}
}

func boardFromQueenCol(queenCol []int)[]string{
	res := make([]string, len(queenCol))
	for i := 0; i < len(queenCol); i ++ {
		byteArr := make([]byte, len(queenCol))
		for k := range byteArr {
			if queenCol[i] == k {
				byteArr[k] = 'Q'
			} else {
				byteArr[k] = '.'
			}
		}
		res[i] = string(byteArr)
	}
	return res
}

func isValid(queenCol []int, row, col int) bool{
	for i := 0; i < row; i ++ {
		if queenCol[i] == col {
			return false
		}
		if abs(row, i) == abs(col, queenCol[i]) {
			return false
		}
 	}
	return true
}

func abs(a, b int)int {
	res := a - b
	if res < 0 {
		res = -res
	}
	return res
}