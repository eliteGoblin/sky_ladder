from collections import defaultdict

class Solution:
    def firstUniqChar(self, s: str) -> int:
        dt = defaultdict(int)
        for c in s:
            dt[c] += 1
        for i, c in enumerate(s):
            if dt[c] == 1:
                return i
        return -1
