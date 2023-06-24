from collections import defaultdict

def contain_dup(dct) -> bool:
    for v in dct.values():
        if v > 1:
            return True
    return False

class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:

        window = defaultdict(int)
        left, right = 0, 0

        res = 1

        while right < len(s):
            window[s[right]] += 1
            right += 1

            print(left, right)

            while left < right:
                if not contain_dup(window):
                    if right - left > res:
                        res = right - left
                    break

                left_c = window[left]
                window[left_c] -= 1
                if window[left_c] == 0:
                    del window[left_c]
                left += 1
                    
        return res
