# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        pseudo_head = ListNode()
        
        while head is not None:
            next_node = head.next
            head.next = pseudo_head.next
            pseudo_head.next = head

            head = next_node
        
        return pseudo_head.next