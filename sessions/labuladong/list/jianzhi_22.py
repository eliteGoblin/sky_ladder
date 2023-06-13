class Solution:
    def getKthFromEnd(self, head: ListNode, k: int) -> ListNode:
        left, right = head, head

        while k > 0 and right is not None:
            right = right.next
            k -= 1
        
        if k > 0:
            return None

        while right is not None:
            left = left.next
            right = right.next
        
        return left