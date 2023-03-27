class Solution:

    def __init__(self):
        self.cache = {} # end potision's sum

    def max_subarray_ends_at(self, nums, index) -> int:
        if index == 0:
            return nums[0]
        if index in self.cache:
            return self.cache[index]
        res = max(
            nums[index], 
            self.max_subarray_ends_at(nums, index - 1) + nums[index]
        )
        self.cache[index] = res
        return res
    
    def maxSubArray(self, nums: List[int]) -> int:
        max_sum = nums[0]

        for i in range(len(nums)):
            sum_end_at_i = self.max_subarray_ends_at(nums, i)
            max_sum = max(max_sum, sum_end_at_i)
        
        return max_sum

        
        