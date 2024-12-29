# Definition for a binary tree node.

from typing import Optional

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    node_max_height: dict[int, int] = {}
    def get_max_height(self, root: Optional[TreeNode]) -> int:
        if root is None:
            return 0
        if id(root) not in self.node_max_height:
            height: int =  1 + max(
                self.get_max_height(root.left),
                self.get_max_height(root.right),
            )
            self.node_max_height[id(root)] = height
        return self.node_max_height[id(root)]

    def diameterOfBinaryTree(self, root: Optional[TreeNode]) -> int:
        if root is None:
            return 0
        height_root: int = self.get_max_height(root)
        return max(
            height_root,
            self.diameterOfBinaryTree(root.left),
            self.diameterOfBinaryTree(root.right),
        )


