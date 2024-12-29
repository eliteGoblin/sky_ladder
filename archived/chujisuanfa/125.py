class Solution:
    def isPalindrome(self, s: str) -> bool:
        strip = ""
        for e in s.lower():
            if e.isalnum():
                strip += e
        
        start, end = 0, len(strip) - 1

        while start < end:
            if strip[start] != strip[end]:
                return False
            start += 1
            end -= 1
        
        return True
    
s = Solution()

print(s.isPalindrome("OP"))