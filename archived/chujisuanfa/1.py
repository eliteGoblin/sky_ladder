from collections import defaultdict

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        dt = defaultdict(list)

        for i, e in enumerate(nums):
            dt[e].append(i)

        for i, e in enumerate(nums):
            pos_arr = dt[target - e]
            if e != target - e:
                if len(pos_arr) > 0:
                    return [i, pos_arr[0]]
            else:
                if len(pos_arr) > 1:
                    return [i, pos_arr[1]]
        
        return [-1. -1]
