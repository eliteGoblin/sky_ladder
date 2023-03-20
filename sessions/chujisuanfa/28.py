class Solution:
    def strStr(self, haystack: str, needle: str) -> int:
        index = -1
        for i in range(len(haystack)):
            s = haystack[i:i+len(needle)]
            if len(s) < len(needle):
                break
            if s == needle:
                index = i
                break
        return index