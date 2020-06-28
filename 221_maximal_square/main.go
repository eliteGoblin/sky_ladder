package _221_maximal_square

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	res := 0
	row, col := len(matrix), len(matrix[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] > '0' {
				maxSquare := findMaxSquare(matrix, i, j)
				if maxSquare > res {
					res = maxSquare
				}
			}
		}
	}
	return res * res
}

func findMaxSquare(matrix [][]byte, i, j int) int {
	max := 1
	row, col := len(matrix), len(matrix[0])
	for size := 2; ; size++ {
		endRow, endCol := i+size, j+size
		if endRow > row || endCol > col {
			return max
		}
		for k := j; k < endCol; k++ {
			if matrix[endRow-1][k] == '0' {
				return max
			}
		}
		for k := i; k < endRow; k++ {
			if matrix[k][endCol-1] == '0' {
				return max
			}
		}
		max++
	}
}
