class Solution:
    def removeNthFromEnd(self, head: ListNode, n: int) -> ListNode:
        pseudo_head = ListNode(next=head)
        left, right = pseudo_head, pseudo_head

        while n + 1 > 0 and right is not None:
            right = right.next
            n -= 1
        
        while right is not None:
            left = left.next
            right = right.next
        
        left.next = left.next.next

        return pseudo_head.next