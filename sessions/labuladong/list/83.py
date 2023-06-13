# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def deleteDuplicatesSinglePointer(self, head: Optional[ListNode]) -> Optional[ListNode]:

        if head is None or head.next is None:
            return head
        
        pre = head
        pre_value = head.val

        cur = head.next.next

        while cur is not None:
            if cur.val != pre_value:
                pre_value = cur.val
                pre.next = cur
            cur = cur.next
        
        return head

class Solution:
    def deleteDuplicates(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if head is None:
            return head
        slow, fast = head, head

        while fast is not None:
            if fast.val != slow.val:
                slow.next = fast
                slow = slow.next
            fast = fast.next
        
        slow.next = None

        return head