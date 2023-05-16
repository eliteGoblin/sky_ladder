# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def removeElements(self, head: Optional[ListNode], val: int) -> Optional[ListNode]:
        pseudo_head = ListNode(next=head)
        pre = pseudo_head
        while pre.next is not None:
            next_node = pre.next
            if next_node.val == val:
                pre.next = pre.next.next
            else:
                pre = pre.next
        return pseudo_head.next
