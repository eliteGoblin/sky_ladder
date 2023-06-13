# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def detectCycle(self, head: ListNode) -> ListNode:
        
        cache = set()
        p = head
        while p is not None:
            if id(p) in cache:
                return p
            cache.add(id(p))

            p = p.next
        
        return None