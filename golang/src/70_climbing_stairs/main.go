package _70_climbing_stairs

func climbStairs(n int) int {
	mp := make(map[int]int)
	return helper(mp, n)
}

func helper(memo map[int]int, n int) int {
	if v, ok := memo[n]; ok {
		return v
	}
	if n == 0 || n == 1 {
		memo[n] = 1
	} else {
		memo[n] = helper(memo, n-1) + helper(memo, n-2)
	}
	return memo[n]
}
