class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        next_fill = 0
        for e in nums:
            if e != val:
                nums[next_fill] = e
                next_fill += 1
        return next_fill