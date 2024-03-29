# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def mergeTwoLists(self, l1: ListNode, l2: ListNode) -> ListNode:

        pseudo_head = ListNode()
        pre = pseudo_head

        while l1 is not None and l2 is not None:
            if l1.val <= l2.val:
                pre.next = l1
                l1 = l1.next
            else:
                pre.next = l2
                l2 = l2.next
            pre = pre.next
        
        if l1 is not None:
            pre.next = l1
        if l2 is not None:
            pre.next = l2
        
        return pseudo_head.next
