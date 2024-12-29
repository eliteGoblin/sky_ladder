class Solution:
    def minOperations(self, s: str) -> int:
        # op to make into 0101... str
        res_01: int = sum(int(int(c) == i%2) for i, c in enumerate(s))