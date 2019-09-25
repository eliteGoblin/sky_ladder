package amzon

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	heights := make([][]int, m)
	for i := range heights {
		heights[i] = make([]int, n)
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i == 0 || matrix[i][j] == '0' {
				heights[i][j] = int(matrix[i][j] - '0')
			} else {
				heights[i][j] = heights[i-1][j] + 1
			}
		}
	}
	maxArea := 0
	for i := range heights {
		area := largestRectangleArea(heights[i])
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

func largestRectangleArea(heights []int) int {
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		if heights[i] > maxArea {
			maxArea = heights[i]
		}
		if i+1 < len(heights) && heights[i] <= heights[i+1] {
			continue
		}
		minHeight := heights[i]
		for j := i - 1; j >= 0; j-- {
			if heights[j] < minHeight {
				minHeight = heights[j]
			}
			area := minHeight * (i - j + 1)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}
