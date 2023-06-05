class Solution:
    def middleNode(self, head: Optional[ListNode]) -> Optional[ListNode]:
        fast, slow = head, head

        while fast is not None and fast.next is not None:
            fast = fast.next.next
            slow = slow.next
        
        return slow