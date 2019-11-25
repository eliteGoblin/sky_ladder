package amazon

import "container/heap"

func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0)
	pq := NewHeap(make([]interface{}, 0), func(x, y interface{}) bool {
		ex := x.([2]int)
		ey := y.([2]int)
		return ex[0] > ey[0]
	})
	for i, v := range nums {
		for pq.Len() > 0 {
			e := pq.Top().([2]int)
			if e[1] <= i-k {
				heap.Pop(pq)
			} else {
				break
			}
		}
		heap.Push(pq, [2]int{v, i})
		if i >= k-1 {
			res = append(res, pq.Top().([2]int)[0])
		}
	}
	return res
}
