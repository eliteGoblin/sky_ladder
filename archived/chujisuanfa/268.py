

class Solution:
    def missingNumber(self, nums: List[int]) -> int:
        nums = sorted(nums)

        start, end = 0, len(nums)
        while start < end:
            mid = start + (end - start) // 2
            if mid == nums[mid]:
                start = mid + 1
            else:
                end = mid
        
        return end
    
class Solution2:
    def missingNumber(self, nums: List[int]) -> int:
        res = 1
        for i, e in enumerate(nums):
            res ^= e
            res ^= i
        res ^= len(nums)
        return res