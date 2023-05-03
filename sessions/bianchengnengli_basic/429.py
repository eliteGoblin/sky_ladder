import queue
"""
# Definition for a Node.
class Node:
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children
"""

class Solution:
    def levelOrder(self, root: 'Node') -> List[List[int]]:
        if root is None:
            return []
        level = queue.Queue()
        level.put(root)

        res = []
        cur_level_res = []
        next_level = queue.Queue()

        while not level.empty():
            node = level.get()
            cur_level_res.append(node.val)
            if node.children is not None:
                for e in node.children:
                    next_level.put(e)

            if level.empty():
                res.append(cur_level_res)
                level, next_level = next_level, level
                cur_level_res = []
        return res


