package amazon

func splitArray(nums []int, m int) int {
	sum := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, len(nums)+1)
	}
	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = intMax
		}
	}
	dp[0][0] = 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= len(nums); j++ {
			for k := i - 1; k < j; k++ {
				val := max(dp[i-1][k], sum[j]-sum[k])
				dp[i][j] = min(dp[i][j], val)
			}
		}
	}
	return dp[m][len(nums)]
}
