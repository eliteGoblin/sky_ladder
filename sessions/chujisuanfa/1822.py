

class Solution:
    def arraySign(self, nums: List[int]) -> int:
        pos = True
        for e in nums:
            if e == 0:
                return 0
            if e < 0:
                pos = not pos
        return 1 if pos else -1