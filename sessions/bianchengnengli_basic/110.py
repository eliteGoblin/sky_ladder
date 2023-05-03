# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

def tree_hight(root) -> int:
    if root is None:
        return 0
    lh = tree_hight(root.left)
    rh = tree_hight(root.right)
    if lh < 0 or rh < 0:
        return -1
    if abs(lh - rh) > 1:
        return -1, False
    return max(lh, rh) + 1
    

class Solution:
    def isBalanced(self, root: Optional[TreeNode]) -> bool:
        if root is None:
            return True
        height = tree_hight(root)
        if height < 0:
            return False
        return True
        