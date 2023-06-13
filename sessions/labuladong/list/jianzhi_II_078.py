# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

import heapq

class PriorityQueue:

    def __init__(self):
        self._pq = []
        self._index = 0

    def push(self, item: ListNode) -> None:
        heapq.heappush(self._pq, [item.val, self._index, item])
        self._index += 1
    
    def pop(self) -> ListNode:
        return heapq.heappop(self._pq)[-1]
    
    def is_empty(self) -> bool:
        return len(self._pq) == 0

class Solution:
    def mergeKLists(self, lists: List[ListNode]) -> ListNode:

        pq = PriorityQueue()

        for head in lists:
            if head is not None:
                pq.push(head)
        
        pseudo_head = ListNode()
        pre = pseudo_head

        while not pq.is_empty():
            node = pq.pop()
            pre.next = node
            pre = pre.next

            if node.next is not None:
                pq.push(node.next)
        
        return pseudo_head.next

