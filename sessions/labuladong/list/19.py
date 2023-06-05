# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

def next_n(nd: ListNode, steps: int) -> ListNode:

    while nd is not None and steps > 0:
        nd = nd.next
        steps -= 1
    
    return nd


class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        pseudo_head = ListNode(next=head)
        
        left = pseudo_head
        right = next_n(left, n+1)

        while right is not None:
            right = right.next
            left = left.next
        
        left.next = left.next.next

        return pseudo_head.next
