class Solution:
    def isMonotonic(self, nums: List[int]) -> bool:
        if len(nums) == 1:
            return True
        is_increasing = 0 # 0 unkown, 1 increasing, 2 decreasing
        for i in range(1, len(nums)):
            diff = nums[i] - nums[i-1]
            if diff == 0:
                continue
            if is_increasing == 0:
                is_increasing = 1 if diff > 0 else -1
                continue
            if is_increasing * diff < 0:
                return False
        return True