from typing import Optional

class Solution1:
    def deleteDuplicates(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if head is None or head.next is None:
            return head
        
        pseudo_head = ListNode()
        pre = pseudo_head

        while head:
            if head.next is not None and head.val == head.next.val:
                while head.next is not None and head.val == head.next.val:
                    head = head.next
                head = head.next
            else:
                pre.next = head
                pre = pre.next
                head = head.next
        
        pre.next = None

        return pseudo_head.next

# recursive
class Solution2:
    def deleteDuplicates(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if head is None or head.next is None:
            return head

        if head.val != head.next.val:
            head.next = self.deleteDuplicates(head.next)
            return head
        
        while head.next and head.val == head.next.val:
            head = head.next
        
        head = head.next
        
        return self.deleteDuplicates(head)
