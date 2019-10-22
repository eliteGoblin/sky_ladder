package amazon

import "sort"

func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}
	ct := make([]int, 26)
	for _, v := range tasks {
		ct[v-'A']++
	}
	sort.Ints(ct)

	mx := ct[25]
	i := 25
	for i >= 0 && ct[i] == mx {
		i--
	}
	return (mx-1)*(n+1) + 25 - i
}
