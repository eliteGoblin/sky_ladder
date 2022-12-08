package _160_intersection_of_two_linked_lists

import "github.com/eliteGoblin/sky_ladder/common"

func getIntersectionNode(headA, headB *common.ListNode) *common.ListNode {
	stkA := common.NewStack(128)
	stkB := common.NewStack(128)

	for headA != nil {
		stkA.Push(headA)
		headA = headA.Next
	}
	for headB != nil {
		stkB.Push(headB)
		headB = headB.Next
	}
	var lastMatch *common.ListNode
	for stkA.Len() > 0 && stkB.Len() > 0 {
		curA, _ := stkA.Top()
		curB, _ := stkB.Top()
		if curA != curB {
			break
		}
		lastMatch = curA.(*common.ListNode)
		stkA.Pop()
		stkB.Pop()
	}
	return lastMatch
}
