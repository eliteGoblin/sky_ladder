from collections import defaultdict

class Solution:
    def firstUniqChar(self, s: str) -> int:
        dct = defaultdict(int)

        for c in s:
            dct[c] += 1
        
        for i, c in enumerate(s):
            if dct[c] == 1:
                return i
        
        return -1