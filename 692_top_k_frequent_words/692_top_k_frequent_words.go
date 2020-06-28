package sixninetwo

import "sort"

func topKFrequent(words []string, k int) []string {
	ct := make(map[string]int)
	for _, v := range words {
		ct[v]++
	}
	bucket := make([][]string, len(words))
	for k, v := range ct {
		bucket[v] = append(bucket[v], k)
	}
	i := len(bucket) - 1
	res := make([]string, 0, k)
	for ; i >= 0; i-- {
		if len(bucket[i]) > 0 {
			sort.Strings(bucket[i])
			res = append(res, bucket[i]...)
		}
		if len(res) >= k {
			break
		}
	}
	return res[0:k]
}
