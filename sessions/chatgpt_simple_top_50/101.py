
def is_two_tree_sym(root1, root2) -> bool:
    if root1 is None and root2 is None:
        return True
    if root1 is None or root2 is None:
        return False
    return (root1.val == root2.val) and is_two_tree_sym(root1.left, root2.right) and is_two_tree_sym(root1.right, root2.left)

class Solution:
    def isSymmetric(self, root: Optional[TreeNode]) -> bool:
        if root is None:
            return True
        return is_two_tree_sym(root.left, root.right)
        