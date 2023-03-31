class Solution:
    def reverseBits(self, n: int) -> int:
        res = 0
        count = 32
        while count > 0:
            res = res * 2 + n % 2
            n //= 2
            count -= 1

        return res