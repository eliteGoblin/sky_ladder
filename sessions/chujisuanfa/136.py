class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        res = 0
        for e in nums:
            res = res ^ e
        return res