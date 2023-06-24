from collections import defaultdict

class Solution:
    def checkInclusion(self, s1: str, s2: str) -> bool:
        s, t = s2, s1
        window, target = defaultdict(int), defaultdict(int)

        for c in t:
            target[c] += 1

        matched_char_count = 0
        left, right = 0, 0

        while right < len(s):
            right_c = s[right]
            right += 1

            if right_c in target and window[right_c] == target[right_c]:
                window[right_c] += 1
                matched_char_count += 1
            
            while left < right and right - left >= len(t):
                if matched_char_count == len(target):
                    return True
                
                left_c = s[left]
                left += 1

                if left_c in target:
                    window[left_c] -= 1
                    if window[left_c] != target[left_c]:
                        matched_char_count -= 1
        
        return False
