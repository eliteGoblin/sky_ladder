# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def getIntersectionNode(self, headA: ListNode, headB: ListNode) -> ListNode:
        p_a, p_b = headA, headB
        a_visited_b, b_visited_a = False, False

        while p_a is not None and p_b is not None:
            if p_a == p_b:
                return p_a
            
            if p_a.next is None and not a_visited_b:
                p_a = headB
                a_visited_b = True
            else:
                p_a = p_a.next
            
            if p_b.next is None and not b_visited_a:
                p_b = headA
                b_visited_a = True
            else:
                p_b = p_b.next
        
        return None