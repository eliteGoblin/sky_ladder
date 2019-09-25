package amazon

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
