package _48_rotate_image

func rotate(matrix [][]int) {
	n := len(matrix)
	if n == 0 {
		return
	}
	// up down switch
	i, j := 0, n-1
	for i < j {
		for k := 0; k < n; k++ {
			matrix[i][k], matrix[j][k] = matrix[j][k], matrix[i][k]
		}
		i++
		j--
	}
	// diag switch
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
