class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        if len(nums) == 0:
            return 0
        
        next_fill = 1

        for cur in range(1, len(nums)):
            if nums[cur] != nums[cur-1]:
                nums[next_fill] = nums[cur]
                next_fill += 1
        return next_fill