class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        if len(nums) == 0:
            return 0
        next_fill_index, check_index = 1, 1
        while check_index < len(nums):
            if nums[check_index] != nums[check_index - 1]:
                nums[next_fill_index] = nums[check_index]
                next_fill_index += 1
            check_index += 1
        return next_fill_index