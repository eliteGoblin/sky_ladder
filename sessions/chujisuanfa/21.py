# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def mergeTwoLists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
        pseudo_head = ListNode()
        pre = pseudo_head

        p1, p2 = list1, list2

        while p1 is not None and p2 is not None:
            if p1.val <= p2.val:
                pre.next = p1
                pre = pre.next
                p1 = p1.next
            else:
                pre.next = p2
                pre = pre.next
                p2 = p2.next
        
        if p1 is not None:
            pre.next = p1
        else:
            pre.next = p2
        
        return pseudo_head.next