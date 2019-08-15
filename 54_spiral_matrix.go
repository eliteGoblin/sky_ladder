package amazon

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	res := make([]int, 0, len(matrix)*len(matrix[0]))
	up, down, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	for {
		for j := left; j <= right; j++ {
			res = append(res, matrix[up][j])
		}
		if up++; up > down {
			break
		}
		for i := up; i <= down; i++ {
			res = append(res, matrix[i][right])
		}
		if right--; right < left {
			break
		}
		for j := right; j >= left; j-- {
			res = append(res, matrix[down][j])
		}
		if down--; down < up {
			break
		}
		for i := down; i >= up; i-- {
			res = append(res, matrix[i][left])
		}
		if left++; left > right {
			break
		}
	}
	return res
}
