
class Solution:
    cache: dict[int, int] # stairs => ways num
    def __init__(self) -> None:
        self.cache = {}
        self.cache[1] = 1
        self.cache[2] = 2

    def climbStairs(self, n: int) -> int:
        if n in self.cache:
            return self.cache[n]
        if n <= 0:
            return 0
        res = self.climbStairs(n-2) + self.climbStairs(n-1)
        self.cache[n] = res
        return res
        

        
