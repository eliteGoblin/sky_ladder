package _957_prison_cells_after_n_days

func prisonAfterNDays(cells []int, N int) []int {
	var init [8]int
	for i := 0; i < 8; i++ {
		init[i] = cells[i]
	}
	mp := make(map[[8]int]int)
	cur := init
	for N > 0 {
		mp[cur] = N
		N--
		cur = next(cur)
		if _, ok := mp[cur]; ok {
			cycle := mp[cur] - N
			N %= cycle
		}
	}
	res := make([]int, 8)
	for i := 0; i < 8; i++ {
		res[i] = cur[i]
	}
	return res
}

func next(cur [8]int) [8]int {
	var res [8]int
	for i := 1; i <= 6; i++ {
		if cur[i-1] == cur[i+1] {
			res[i] = 1
		} else {
			res[i] = 0
		}
	}
	return res
}

// 以下解法挂在最后几个case:  [1,0,0,0,1,0,0,1], 99 WA, 但98能过, cycle为14, cur为84
func prisonAfterNDaysAWrong(cells []int, N int) []int {
	n := N
	var init [8]int
	for i := 0; i < 8; i++ {
		init[i] = cells[i]
	}
	mp := make(map[[8]int]int)
	cur := init

	var cycle int
	for n > 0 {
		if _, ok := mp[cur]; ok {
			cycle = mp[cur] - n
			break
		}
		mp[cur] = n
		cur = next(cur)
		N--
	}
	N %= cycle
	for N > 0 {
		cur = next(init)
		N--
	}
	res := make([]int, 8)
	for i := 0; i < 8; i++ {
		res[i] = cur[i]
	}
	return res
}

// 错误的解法

func prisonAfterNDaysWrong(cells []int, N int) []int {
	var init [8]int
	for i := 0; i < 8; i++ {
		init[i] = cells[i]
	}
	mp := make(map[[8]int]struct{})
	cur := init
	for N > 0 {
		if _, ok := mp[cur]; ok {
			N %= len(mp) //这里以为size就是cycle, 是错误的，因为init不一定是周期的起点，可能是半中间. 调试发现len(mp) == 15, 其实已经包括一个完整周期再加上额外的一些值了
		}
		mp[cur] = struct{}{}
		cur = next(cur)
		N--
	}
	res := make([]int, 8)
	for i := 0; i < 8; i++ {
		res[i] = cur[i]
	}
	return res
}
