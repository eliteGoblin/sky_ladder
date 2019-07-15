package amazon

type ListNode struct {
	Val  int
	Next *ListNode
}

func max(a... int)int {
	m := a[0]
	for i := 1; i < len(a); i ++ {
		if a[i] > m {
			m = a[i]
		}
	}
	return m
}

func min(a... int)int {
	m := a[0]
	for i := 1; i < len(a); i ++ {
		if a[i] < m {
			m = a[i]
		}
	}
	return m
}