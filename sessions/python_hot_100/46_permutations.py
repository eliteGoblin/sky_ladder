from typing import Optional, Any, Sequence, List

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right



class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        if len(nums) == 0:
            return []
        if len(nums) == 1:
            return [nums[0]]
        next_result = self.permute(nums[1:])
        result: List[List[int]] = []
        for lst in next_result:
            for i in range(0, len(lst)+1):
                new_lst = lst.copy().insert(i)
                result.append(new_lst)
        return result
