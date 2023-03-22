# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
def next_n_node(node, n):
    while n > 0:
        node = node.next
        n -= 1
    return node

class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        pseudo_head = ListNode(val=0, next=head)
        # right pointer move last N + 1 step
        left, right = pseudo_head, next_n_node(pseudo_head, n + 1)

        while right is not None:
            left = next_n_node(left, 1)
            right = next_n_node(right, 1)
        
        left.next = left.next.next

        return pseudo_head.next
