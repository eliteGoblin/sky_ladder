class Solution:
    def lengthOfLastWord(self, s: str) -> int:
        reverse_start = len(s) - 1
        while reverse_start >= 0 and s[reverse_start] == " ":
            reverse_start -= 1
        reverse_last = reverse_start - 1
        
        while reverse_last >= 0 and s[reverse_last] != " ":
            reverse_last -= 1
        
        return reverse_start - reverse_last

        