from header import *

class Solution:
    def flatten(self, root: Optional[TreeNode]) -> None:
        """
        Do not return anything, modify root in-place instead.
        """
        if root is None:
            return
        self.flatten(root.left)
        self.flatten(root.right)

        left_right_most = root.left
        if left_right_most is not None:
            while left_right_most.right is not None:
                left_right_most = left_right_most.right
            left_right_most.right = root.right
            root.right = root.left
            root.left = None
        