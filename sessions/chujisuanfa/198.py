

class Solution:
    def __init__(self):
        self.cache = {}

    def rob_house(self, nums, index) -> int:
        if index < 0:
            return 0
        if index in self.cache:
            return self.cache[index]
        
        res = nums[index] + max(
            self.rob_house(nums, index - 2),
            self.rob_house(nums, index - 3)
        )

        self.cache[index] = res

        return res

    def rob(self, nums: List[int]) -> int:
        if len(nums) == 1:
            return nums[0]
        
        res = max(
            self.rob_house(nums, len(nums) - 1),
            self.rob_house(nums, len(nums) - 2)
        )

        return res