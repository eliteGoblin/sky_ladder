from collections import deque

# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def levelOrderBottom(self, root: Optional[TreeNode]) -> List[List[int]]:
        if root is None:
            return []
        
        res = []
        q = deque()
        q.appendleft(root)

        cur_level = []
        next_level = deque()

        while len(q) > 0:
            node = q.pop()
            cur_level.append(node.val)

            if node.left is not None:
                next_level.appendleft(node.left)
            if node.right is not None:
                next_level.appendleft(node.right)
            
            if len(q) == 0:
                res.append(cur_level)
                cur_level = []
                q, next_level = next_level, q

        return res[::-1]

