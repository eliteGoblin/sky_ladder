package amazon

func coinChange(coins []int, amount int) int {
	mp := make(map[int]int, 0)
	return coinChangeMemo(coins, amount, mp)
}

func coinChangeMemo(coins []int, amount int, mp map[int]int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	res := intMax
	if ct, ok := mp[amount]; ok {
		return ct
	}
	for _, v := range coins {
		ct := coinChangeMemo(coins, amount-v, mp)
		if ct >= 0 {
			if res > ct+1 {
				res = ct + 1
			}
		}
	}
	if res == intMax {
		res = -1
	}
	mp[amount] = res
	return res
}
