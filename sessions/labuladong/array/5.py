class Solution:
    def longestPalindrome(self, s: str) -> str:
        if len(s) <= 1:
            return s

        l_res, r_res = 0, 0

        def search(l, r):
            nonlocal l_res
            nonlocal r_res
            # l <= r
            while l >= 0 and r < len(s):
                if s[l] != s[r]:
                    return
                cur_len = r - l + 1
                if cur_len > (r_res - l_res + 1):
                    l_res, r_res = l, r

                l -= 1
                r += 1
        
        for i in range(len(s)):
            search(i, i)
            search(i, i + 1)

        return s[l_res:r_res+1]