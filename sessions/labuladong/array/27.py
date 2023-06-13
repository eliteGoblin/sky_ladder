class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        slow, fast = 0, 0
        while slow < len(nums):
            if nums[slow] == val:
                continue
            slow += 1
        fast = slow

        while fast < len(nums):
            if val != nums[fast]:
                slow += 1
                nums[slow] = nums[fast]
            fast += 1

        return slow+1