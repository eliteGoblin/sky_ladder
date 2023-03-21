# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def deleteNode(self, node):
        """
        :type node: ListNode
        :rtype: void Do not return anything, modify node in-place instead.
        """
        # node.next not none
        while node.next is not None:
            node.val = node.next.val
            if node.next.next is None:
                # last[-2]
                node.next = None
                break
            node = node.next