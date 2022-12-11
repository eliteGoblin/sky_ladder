from typing import Optional, Any, Sequence, List

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right



class Solution:
    def count_1(self, i: int) -> int:
        res = 0
        while i > 0:
            res += i % 2
            i = i // 2
        return res
    def countBits(self, n: int) -> List[int]:
        res = [0] * ( n + 1) 
        for i in range(0, len(res)):
            res[i] = self.count_1(i)
        return res
