package amazon

func reverseListIterative(head *ListNode) *ListNode {
	pseudoHead := &ListNode{}
	var next *ListNode
	for p := head; p != nil; {
		next = p.Next
		p.Next = pseudoHead.Next
		pseudoHead.Next = p
		p = next
	}
	return pseudoHead.Next
}

func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseListRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
