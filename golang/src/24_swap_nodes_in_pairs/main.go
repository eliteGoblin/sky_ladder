package _24_swap_nodes_in_pairs

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	sec := head.Next
	next := swapPairs(head.Next.Next)
	head.Next = next
	sec.Next = head
	return sec
}
