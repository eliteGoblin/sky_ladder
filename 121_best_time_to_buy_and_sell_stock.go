package amazon

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	m := 0
	preMin := prices[0]
	for i := 1; i < len(prices); i++ {
		m = max(m, prices[i]-preMin)
		preMin = min(prices[i], preMin)
	}
	return m
}
