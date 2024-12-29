class Solution:
    def reverse(self, x: int) -> int:
        is_neg: bool = False
        if x < 0:
            is_neg = True
            x = -x
        res = 0
        while x > 0:
            res = res * 10 + x % 10
            x = x // 10

        if is_neg:
            res = -res
        if -2**31 <= res <= 2**31 - 1:
            return res
        return 0