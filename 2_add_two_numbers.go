package amazon

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	carry := 0
	head := &ListNode{
		Next: nil,
	}
	pre := head
	for l1 != nil && l2 != nil {
		value := l1.Val + l2.Val + carry
		pre.Next = &ListNode{
			Val: value % 10,
		}
		pre = pre.Next
		carry = value / 10
		l1 = l1.Next
		l2 = l2.Next
	}
	left := l1
	if left == nil {
		left = l2
	}
	for left != nil {
		value := left.Val + carry
		pre.Next = &ListNode{
			Val: value % 10,
		}
		pre = pre.Next
		carry = value / 10
		left = left.Next
	}
	if carry > 0 {
		pre.Next = &ListNode{
			Val: carry,
		}
	}
	return head.Next
}
