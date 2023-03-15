class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        next_fill_index, peek_index = 0, 0
        while peek_index < len(nums):
            if nums[peek_index] != 0:
                nums[next_fill_index] = nums[peek_index]
                next_fill_index += 1
            peek_index += 1
        
        for i in range(next_fill_index, len(nums)):
            nums[i] = 0
            