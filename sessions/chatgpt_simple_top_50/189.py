def reverse_inplace(nums, start, last):
    while start < last:
        nums[start], nums[last] = nums[last], nums[start]
        start += 1
        last -= 1

class Solution:
    def rotate(self, nums: List[int], k: int) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        k = k % len(nums)
        reverse_inplace(nums, 0, len(nums)-1)
        reverse_inplace(nums, 0, k-1)
        reverse_inplace(nums, k, len(nums)-1)


