def count_5_times(n: int):
    res = 0
    while n % 5 == 0:
        res += 1
        n = n / 5
    return res

class Solution:
    def trailingZeroes(self, n: int) -> int:
        res = 0
        for i in range(1, n+1):
            res += count_5_times(i)
        return res