# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
import queue

class Solution:
    def levelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        if root is None:
            return []
        
        q_cur, q_next = queue.Queue(), queue.Queue()
        q_cur.put(root)

        res = []
        res_cur = []

        while not q_cur.empty():
            node = q_cur.get()
            res_cur.append(node.val)

            if node.left is not None:
                q_next.put(node.left)
            if node.right is not None:
                q_next.put(node.right)
    
            if q_cur.empty(): # move to next layer
                res.append(res_cur)
                q_cur = q_next
                q_next = queue.Queue()
                res_cur = []
        
        return res