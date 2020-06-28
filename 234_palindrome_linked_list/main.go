package _234_palindrome_linked_list

import "github.com/eliteGoblin/sky_ladder/common"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *common.ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	p := slow.Next
	slow.Next = nil
	for p != nil {
		next := p.Next
		p.Next = slow.Next
		slow.Next = p
		p = next
	}
	p = slow.Next
	for p != nil {
		if head.Val != p.Val {
			return false
		}
		head = head.Next
		p = p.Next
	}
	return true
}
