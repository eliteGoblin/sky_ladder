# Definition for singly-linked list.

import heapq

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def partition(self, head: Optional[ListNode], x: int) -> Optional[ListNode]:
        lt_pseudo_head = ListNode()
        p_lt = lt_pseudo_head
        ge_pseudo_head = ListNode()
        p_ge = ge_pseudo_head

        cur = head

        while cur is not None:
            if cur.val < x:
                p_lt.next = cur
                p_lt = p_lt.next
            else:
                p_ge.next = cur
                p_ge = p_ge.next

            cur = cur.next
            
        p_ge.next = None
        p_lt.next = ge_pseudo_head.next

        return lt_pseudo_head.next