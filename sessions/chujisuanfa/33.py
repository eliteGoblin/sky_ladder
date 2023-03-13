class Solution:
    def rotate(self, nums: List[int], k: int) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        k = k % len(nums)
        tmp = nums[:k]
        for i in range(len(nums)-1, k-1, -1):
            nums[i-k] = nums[i]
        nums[len-k:] = tmp