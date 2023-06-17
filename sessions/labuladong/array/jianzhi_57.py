class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        l, r = 0, len(nums) - 1

        while l < r:
            sm = nums[l] + nums[r]
            if sm == target:
                return [nums[l], nums[r]]
            elif sm < target:
                l += 1
            else:
                r -= 1
        
        return [-1, -1]