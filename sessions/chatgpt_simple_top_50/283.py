class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        next_fill, cur = 0, 0

        while cur < len(nums):
            if nums[cur] != 0:
                nums[next_fill] = nums[cur]
                next_fill += 1
            cur += 1
        
        for i in range(next_fill, len(nums)):
            nums[i] = 0
        