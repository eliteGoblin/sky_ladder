class Solution:
    def arraySign(self, nums: List[int]) -> int:
        res: int = 1
        for e in nums:
            if e == 0:
                return 0
            if e < 0:
                res *= -1
        return res
