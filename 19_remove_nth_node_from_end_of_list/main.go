package _19_remove_nth_node_from_end_of_list

import "github.com/eliteGoblin/sky_ladder/common"

func removeNthFromEnd(head *common.ListNode, n int) *common.ListNode {
	if n == 0 || head == nil {
		return head
	}
	pseudoHead := &common.ListNode{Next: head}
	probe := pseudoHead
	var i int
	for i = 0; i <= n; i++ {
		probe = probe.Next
	}
	var cur *common.ListNode
	for cur = pseudoHead; probe != nil; {
		cur = cur.Next
		probe = probe.Next
	}
	cur.Next = cur.Next.Next
	return pseudoHead.Next
}
