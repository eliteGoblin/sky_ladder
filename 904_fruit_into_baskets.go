package amazon

func totalFruit(tree []int) int {
	res := 0
	mpBaskets := make(map[int]int)
	start := 0
	for i, v := range tree {
		mpBaskets[v]++
		for len(mpBaskets) > 2 {
			mpBaskets[tree[start]]--
			if mpBaskets[tree[start]] == 0 {
				delete(mpBaskets, tree[start])
			}
			start++
		}
		if i-start+1 > res {
			res = i - start + 1
		}
	}
	return res
}
