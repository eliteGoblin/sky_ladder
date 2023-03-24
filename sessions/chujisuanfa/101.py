# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

def is_symmetric(root1, root2) -> bool:
    if root1 is None and root2 is None:
        return True
    if root1 is not None and root2 is not None:
        if root1.val != root2.val:
            return False
        return is_symmetric(root1.left, root2.right) and is_symmetric(root1.right, root2.left)
    return False

class Solution:
    def isSymmetric(self, root: Optional[TreeNode]) -> bool:
        return is_symmetric(root.left, root.right)