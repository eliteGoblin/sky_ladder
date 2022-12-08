from typing import Optional, Any, Sequence, List

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right



class Solution:
    def mergeTrees(self, root1: Optional[TreeNode], root2: Optional[TreeNode]) -> Optional[TreeNode]:
        if root1 is not None and root2 is not None:
            new_root = TreeNode(val=root1.val+root2.val)
            new_root.left = self.mergeTrees(root1.left, root2.left)
            new_root.right = self.mergeTrees(root1.right, root2.right)
            return new_root
        return root1 if root1 is not None else root2
