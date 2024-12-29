from typing import Optional, Any, Sequence, List
from header import TreeNode

# Definition for a binary tree node.

class Solution:
    def inorder_recursive(self, root: Optional[TreeNode], path: List[int]):
        if root is None:
            return
        self.inorder_recursive(root.left, path)
        path.append(root.val)
        self.inorder_recursive(root.right, path)

    def inorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        path = []
        self.inorder_recursive(root, path)
        return path
