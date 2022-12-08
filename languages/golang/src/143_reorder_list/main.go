package _143_reorder_list

import "github.com/eliteGoblin/sky_ladder/common"

func reorderList(head *common.ListNode) {
	if head == nil {
		return
	}
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	secList := slow.Next
	slow.Next = nil

	pseudoHead := &common.ListNode{}
	for p := secList; p != nil; {
		next := p.Next
		p.Next = pseudoHead.Next
		pseudoHead.Next = p
		p = next
	}
	// merge
	p1, p2 := head, pseudoHead.Next
	for p2 != nil {
		p1Next := p1.Next
		p2Next := p2.Next
		p1.Next = p2
		p2.Next = p1Next
		p1 = p1Next
		p2 = p2Next
	}
}
