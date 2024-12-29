from collections import defaultdict

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        dt = defaultdict(list)
        for i, e in enumerate(nums):
            dt[e].append(i)
        
        for i, e in enumerate(nums):
            op2 = target - e
            if op2 in dt:
                for j in dt[op2]:
                    if j != i:
                        return [i, j]
        
        return [-1, -1]