from collections import defaultdict

class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        window = defaultdict(int)

        left, right = 0, 0

        non_dup_max_substr_len = 0

        while right < len(s):
            right_c = s[right]
            window[right_c] += 1
            right += 1

            while left < right and window[right_c] > 1:
                # shrink if contain dup
                # no dup, update res if larger
                left_c = s[left]
                window[left_c] -= 1
                if window[left_c] == 0:
                    del window[left_c]
                left += 1
                
            non_dup_max_substr_len = max(non_dup_max_substr_len, right - left) 
                
        return non_dup_max_substr_len
