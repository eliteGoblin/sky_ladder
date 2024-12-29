class Solution:
    def climbStairs(self, n: int) -> int:
        count = [0] * n
        count[0], count[1] = 1, 2
        for i in range(2, n):
            count[i] = count[i-1] + count[i-2]
        return count[-1]