package amazon

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	mid := len(lists) / 2
	l1 := mergeKLists(lists[:mid])
	l2 := mergeKLists(lists[mid:])
	pseudoHead := &ListNode{}
	pre := pseudoHead
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			pre.Next = l1
			pre = l1
			l1 = l1.Next
		} else {
			pre.Next = l2
			pre = l2
			l2 = l2.Next
		}
	}
	if l1 != nil {
		pre.Next = l1
	} else {
		pre.Next = l2
	}
	return pseudoHead.Next
}
