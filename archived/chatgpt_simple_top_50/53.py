class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        end_at = [0] * len(nums)
        end_at[0] = nums[0]
        for i in range(1, len(nums)):
            end_at[i] = max(nums[i], end_at[i-1] + nums[i])
        return max(end_at)
