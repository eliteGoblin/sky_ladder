# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

MAX_VALUE = 2 ** 31
MIN_VALUE = - (2 ** 31) - 1

def isValidBSTRecur(root, min_value, max_value) -> bool:
    if root is None:
        return True
    if not (min_value < root.val < max_value):
        return False
    return isValidBSTRecur(root.left, min_value, root.val) and isValidBSTRecur(root.right, root.val, max_value)


class Solution:
    def isValidBST(self, root: Optional[TreeNode]) -> bool:
        return isValidBSTRecur(root, MIN_VALUE, MAX_VALUE)
