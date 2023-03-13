# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

def sum_recur(is_left: bool, root: Optional[TreeNode]) -> int:
    if root is None:
        return 0
    if is_left and root.left is None and root.right is None:
        return root.val
    return sum_recur(True, root.left) + sum_reur(False, root.right)
    
class Solution:
    def sumOfLeftLeaves(self, root: Optional[TreeNode]) -> int:
        return sum_recur(False, root)
