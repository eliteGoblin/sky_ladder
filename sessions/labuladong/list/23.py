# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

import heapq


class PriorityQueueMinHeap:
    def __init__(self) -> None:
        self._pq = []
        self._index = 0

    def push(self, node: ListNode) -> None:
        heapq.heappush(self._pq, [node.val, self._index, node])
        self._index += 1
    
    def pop(self) -> ListNode:
        if len(self._pq) == 0:
            raise "empty pq cannot pop"
        return heapq.heappop(self._pq)[-1]

    def empty(self) -> bool:
        return len(self._pq) == 0

class Solution:
    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        pq = []
        index = 0

        pseudo_head = ListNode()
        pre = pseudo_head

        pq = PriorityQueueMinHeap()

        for list_head in lists:
            if list_head is None:
                continue
            pq.push(list_head)
        
        while not pq.empty():
            cur_node = pq.pop()

            pre.next = cur_node
            pre = pre.next

            if cur_node.next is not None:
                pq.push(cur_node.next)
        
        return pseudo_head.next
