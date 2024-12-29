class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        next_fill_index, cur_index = 0, 0
        for cur_index in range(len(nums)):
            if nums[cur_index] != 0:
                nums[next_fill_index] = nums[cur_index]
                next_fill_index += 1

        for i in range(next_fill_index, len(nums)):
            nums[i] = 0

class Solution:
    def middleNode(self, head: Optional[ListNode]) -> Optional[ListNode]:
        