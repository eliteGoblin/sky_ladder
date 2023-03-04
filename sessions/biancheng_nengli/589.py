"""
# Definition for a Node.
class Node:
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children
"""

def preorder_recur(root, res):
    if root is None:
        return
    res.append(root.val)
    for nd in root.children:
        preorder_recur(nd, res)
    
class Solution:
    def preorder(self, root: 'Node') -> List[int]:
        res = []
        preorder_recur(root, res)
        return res