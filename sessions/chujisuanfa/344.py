class Solution:
    def reverseString(self, s: List[str]) -> None:
        """
        Do not return anything, modify s in-place instead.
        """
        start, last = 0, len(s) - 1
        while start < last:
            s[start], s[last] = s[last], s[start]
            start += 1
            last -= 1
        return s