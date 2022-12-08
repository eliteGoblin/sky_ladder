package _203_remove_linked_list_elements

import "github.com/eliteGoblin/sky_ladder/common"

func removeElements(head *common.ListNode, val int) *common.ListNode {
	pseudoHead := &common.ListNode{
		Next: head,
	}
	pre := pseudoHead
	for pre.Next != nil {
		if pre.Next.Val == val {
			pre.Next = pre.Next.Next
		} else {
			pre = pre.Next
		}
	}
	return pseudoHead.Next
}
