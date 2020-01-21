package amazon

import "container/heap"

func compare(a, b interface{}) bool {
	return a.(string) < b.(string)
}

func topKFrequent692(words []string, k int) []string {
	mp := make(map[string]int)
	for _, v := range words {
		mp[v]++
	}
	bucket := make([]*myHeap, len(words))
	for i := range bucket {
		bucket[i] = NewHeap([]interface{}{}, compare)
	}
	for k, v := range mp {
		heap.Push(bucket[v-1], k)
	}
	res := make([]string, 0, len(words))
	for i := len(words) - 1; i >= 0; i-- {
		for bucket[i].Len() > 0 {
			res = append(res, heap.Pop(bucket[i]).(string))
			k--
			if k <= 0 {
				return res
			}
		}
	}
	return res
}
