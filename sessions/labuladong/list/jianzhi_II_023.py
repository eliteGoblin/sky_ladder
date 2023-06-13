# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def getIntersectionNode(self, headA: ListNode, headB: ListNode) -> ListNode:
        cache = set()
        while headA is not None:
            cache.add(id(headA))
            headA = headA.next
        
        while headB is not None:
            if id(headB) in cache:
                return headB
            headB = headB.next
        
        return None