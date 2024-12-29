# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

def is_sub_path_recur(head, root):

    if head is None:
        return True
    if root is None:
        return False
    if head.val != root.val:
        return False
    
    l_res = is_sub_path_recur(head.next, root.left)
    r_res = is_sub_path_recur(head.next, root.right)
    return l_res or r_res

class Solution:

    def isSubPath(self, head: Optional[ListNode], root: Optional[TreeNode]) -> bool:
        if is_sub_path_recur(head, root):
            return True
        lc, rc = False, False
        if root.left is not None:
            lc = self.isSubPath(head, root.left)
        if root.right is not None:
            rc = self.isSubPath(head, root.right)
        return lc or rc