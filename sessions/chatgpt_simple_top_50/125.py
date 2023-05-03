class Solution:
    def isPalindrome(self, s: str) -> bool:
        start, last = 0, len(s) - 1
        while start < last:
            if not("a" <= s[start].lower() <= "z"):
               start += 1
               continue
            if not("a" <= s[last].lower() <= "z"):
                last -= 1
                continue
            print("here", s[start].lower(), s[last].lower())
            if s[start].lower() != s[last].lower():
                return False
            start += 1
            last -= 1
        return True

s = Solution()

print(s.isPalindrome("OP"))