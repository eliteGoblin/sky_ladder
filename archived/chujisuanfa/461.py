class Solution:
    def hammingDistance(self, x: int, y: int) -> int:
        xor = x ^ y

        res = 0
        while xor > 0:
            if xor % 2 == 1:
                res += 1
            xor = xor // 2
        
        return res