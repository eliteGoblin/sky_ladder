from collections import defaultdict

def dict1_include_dict2(dct1, dct2) -> bool:
    for k in dct2:
        if k not in dct1:
            return False
        if dct1[k] < dct2[k]:
            return False
    return True

def compare_window_target(window, target) -> int:
    res1 = dict1_include_dict2(window, target)
    res2 = dict1_include_dict2(target, window)
    if res1 and res2:
        return 0
    if res1 and not res2:
        return 1
    return -1

class Solution:
    def findAnagrams(self, s: str, p: str) -> List[int]:

        window, target = defaultdict(int), defaultdict(int)
        for c in p:
            target[c] += 1
        
        left, right = 0, 0
        res = []

        while right < len(s):
            window[s[right]] += 1
            right += 1

            while left < right:
                cmp_res = compare_window_target(window, target)
                if cmp_res < 0:
                    break
                if cmp_res == 0:
                    res.append(left)
                
                left_c = s[left]
                window[left_c] -= 1
                if window[left_c] == 0:
                    del window[left_c]
                left += 1

        return res