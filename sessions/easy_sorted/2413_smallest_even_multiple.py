class Solution:
    def smallestEvenMultiple(self, n: int) -> int:
        p = n
        while p <= 2 * n:
            if p % 2 == 0 and p % n == 0:
                return p
            p += 1
        return -1