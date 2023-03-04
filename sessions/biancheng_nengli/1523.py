class Solution:
    def countOdds(self, low: int, high: int) -> int:
        res: int = 0
        if low % 2 == 0:
            low += 1
        if high % 2 == 0:
            high -= 1
        if low > high:
            return 0
        return (high - low) // 2 + 1

s = Solution()

print(s.countOdds(3, 7))