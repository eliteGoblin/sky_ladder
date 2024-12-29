class Solution:
    def toLowerCase(self, s: str) -> str:
        res = ""
        for c in s:
            if ord('A') <= ord(c) <= ord('Z'):
                res += chr(ord(c) + 32)
            else:
                res += c
        return res