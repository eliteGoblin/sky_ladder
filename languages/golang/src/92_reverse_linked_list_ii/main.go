package _92_reverse_linked_list_ii

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	pseudoHead := &ListNode{
		Next: head,
	}
	pre := pseudoHead
	for i := 1; i < m; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := m; i < n; i++ {
		t := cur.Next
		cur.Next = cur.Next.Next
		t.Next = pre.Next
		pre.Next = t
	}
	return pseudoHead.Next
}

/*
Reverse: 固定pre和tail, 按顺序插入:新来的节点总是插在pre后面
因此本题中pre和cur都固定不动:后面节点陆续插入pre-->cur区间中
*/
