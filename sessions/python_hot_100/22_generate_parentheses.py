from typing import Optional, Any, Sequence, List

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right



class Solution:
    def generate(self, left_count, right_count, current_path, total):
        if left_count < 0 or right_count < 0 or left_count > right_count:
            return
        if left_count == 0 and right_count == 0:
            total.append(current_path)
            return
        self.generate(left_count-1, right_count, current_path + "(", total)
        self.generate(left_count, right_count-1, current_path+")", total)
        
    def generateParenthesis(self, n: int) -> List[str]:
        total = []
        self.generate(n, n, "", total)
        return total