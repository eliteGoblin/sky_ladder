package _304_range_sum_query_2d_immutable

type NumMatrix struct {
	dp [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return NumMatrix{}
	}
	dp := make([][]int, len(matrix)+1)
	for i := range dp {
		dp[i] = make([]int, len(matrix[0])+1)
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	return NumMatrix{
		dp: dp,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if row2+1 >= len(this.dp) || col2+1 >= len(this.dp[0]) {
		return 0
	}
	dp := this.dp
	return dp[row2+1][col2+1] - dp[row2+1][col1] - dp[row1][col2+1] + dp[row1][col1]
}
