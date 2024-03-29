from collections import defaultdict

def window_cover_target(window, target) -> bool:
    for k in target:
        if k not in window:
            return False
        if target[k] > window[k]:
            return False
    return True

class Solution:
    def minWindow(self, s: str, t: str) -> str:
        window = defaultdict(int)
        left, right = 0, 0 # [left, right) is window

        cover = False
        min_window_str = s

        target = defaultdict(int)
        for c in t:
            target[c] += 1

        while right < len(s):
            window[s[right]] += 1
            right += 1

            while left < right and window_cover_target(window, target):
                c_left = s[left]
                cover = True
                if right - left < len(min_window_str):
                    min_window_str = s[left:right]

                window[c_left] -= 1
                left += 1

                if window[c_left] == 0:
                    del window[c_left]
        
        if not cover:
            return ""
        return min_window_str
    
class Solution:
    def minWindow(self, s: str, t: str) -> str:
        target = defaultdict(int)
        for c in t:
            target[c] += 1
        window = defaultdict(int)

        left, right = 0, 0
        min_cover_str = ""

        while right < len(s):
            right_c = s[right]
            window[right_c] += 1
            right += 1

            while left < right:
                if not window_cover_target(window=window, target=target):
                    break
                if min_cover_str == "" or right - left < len(min_cover_str):
                    min_cover_str = s[left:right]

                left_c = s[left]
                window[left_c] -= 1
                if window[left_c] == 0:
                    del window[left_c]
                left += 1
        
        return min_cover_str
    
