from typing import List

class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        result: List[List[int]] = []
        self.get_subsets(nums, [], result)
        return result
    def get_subsets(self, nums: List[int], path: List[int], final_result: List[List[int]]):
        if len(nums) == 0:
            final_result.append(path)
            return
        # not take nums[0]
        self.get_subsets(nums[1:], path.copy(), final_result)
        # take nums[0]
        new_path_take = path.copy()
        new_path_take.append(nums[0])
        self.get_subsets(nums[1:], new_path_take, final_result)