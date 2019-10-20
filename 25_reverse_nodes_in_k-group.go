package amazon

func reverseKGroup(head *ListNode, k int) *ListNode {
	end := head
	for i := 0; i < k; i++ {
		if end == nil {
			return head
		}
		end = end.Next
	}
	newHead := reverse(head, end)
	head.Next = reverseKGroup(end, k)
	return newHead
}

func reverse(head *ListNode, end *ListNode) *ListNode {
	pseudoHead := &ListNode{}
	for p := head; p != end; {
		next := p.Next
		p.Next = pseudoHead.Next
		pseudoHead.Next = p
		p = next
	}
	return pseudoHead.Next
}
