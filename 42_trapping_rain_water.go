package amazon


func trap(height []int) int {
	res := 0
	dp := make([]int, len(height))
	maxValue := 0
	for i := 0; i < len(dp);i ++ {
		dp[i] = maxValue
		maxValue = max(maxValue, height[i])
	}
	maxValue = 0
	for i := len(dp) - 1; i >= 0; i -- {
		dp[i] = min(dp[i], maxValue)
		maxValue = max(maxValue, height[i])
		if dp[i] > height[i] {
			res += dp[i] - height[i]
		}
	}
    return res
}