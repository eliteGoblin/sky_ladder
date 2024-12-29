class Solution:
    def titleToNumber(self, columnTitle: str) -> int:
        res = 0
        for i in range(len(str)-1, -1. -1):
            res = res * 26 + ord(columnTitle[i]) - ord('A') + 1
        return res