# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class SolutionHash:
    def getIntersectionNode(self, headA: ListNode, headB: ListNode) -> Optional[ListNode]:
        cache: dict[int, ListNode] = set()

        while headA is not None:
            cache[id(headA)] = headA
            headA = headA.next
        
        while headB is not None:
            if id(headB) in cache:
                return cache[id(headB)]
            headB = headB.next
        
        return None

def get_len(head: ListNode) -> int:
    res = 0
    while head is not None:
        res += 1
        head = head.next
    return res

def next_n_steps(head: ListNode, steps: int) -> ListNode:
    while steps > 0:
        head = head.next
        steps -= 1
    return head

class Solution:
    def getIntersectionNode(self, headA: ListNode, headB: ListNode) -> Optional[ListNode]:
        len_a = headA
        len_b = headB

        if len_a < len_b:
            headB = next_n_steps(headB, len_b - len_a)
        else:
            headA = next_n_steps(headA, len_a - len_b)
        
        while headA is not None and headB is not None:
            if headA == headB:
                return headA
            headA = headA.next
            headB = headB.next
        
        return None