from header import *

class Solution:
    def find_left_most(self, root):
        left_most = root
        while left_most.left is not None:
            left_most = left_most.left
        return left_most.val

    def convert(self, root, value):
        if root is None:
            return None

        if root.right is not None:
            self.convert(root.right, value)
            root.val += self.find_left_most(root.right) # error
        else:
            root.val += value

        if root.left is not None:
            self.convert(root.left, root.val)

        return root

    def convertBST(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
        self.convert(root, 0)
        return root