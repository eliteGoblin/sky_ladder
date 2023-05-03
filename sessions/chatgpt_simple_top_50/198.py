class Solution:
    def __init__(self) -> None:
        self.cache = {}

    def get_at(self, nums, i) -> int:
        if i < 0:
            return 0
        if i == 0:
            self.cache[0] = nums[0]
        if i == 1:
            self.cache[1] = nums[1]
        
        if i in self.cache:
            return self.cache[i]
        
        self.cache[i] = nums[i] + max(
            self.get_at(nums, i - 2),
            self.get_at(nums, i - 3)
        )
        
        return self.cache[i]


    def rob(self, nums: List[int]) -> int:
        if len(nums) == 0:
            return 0
        if len(nums) == 1:
            return nums[0]
        
        return max(
            self.get_at(nums, len(nums)-1),
            self.get_at(nums, len(nums)-2)
        )