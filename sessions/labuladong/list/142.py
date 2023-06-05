# Definition for singly-linked list.
from typing import Optional


class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class SolutionCache:
    def detectCycle(self, head: Optional[ListNode]) -> Optional[ListNode]:
        cache: dict[int, ListNode] = {}
        if head is None:
            return None
        
        cur = head

        while cur.next is not None:
            if id(cur.next) == id(head):
                return head
            if id(cur.next) in cache:
                return cur.next
            cache[id(cur.next)] = id(cur)
            cur = cur.next
        
        return None


class Solution:
    def detectCycle(self, head: Optional[ListNode]) -> Optional[ListNode]:
        slow, fast = head, head

        cycle = False
        while fast is not None and fast.next is not None:
            fast = fast.next.next
            slow = slow.next
            if fast == slow:
                cycle = True
                break
        
        if not cycle:
            return None
        
        slow = head
        while slow != fast:
            slow = slow.next
            fast = fast.next
        
        return slow