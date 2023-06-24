from collections import defaultdict

def dict1_comp_dict2(dict1, dict2) -> int:
    # 0 means euqal, -1 keys in dict2 not in dict1; 1: keys in dict2 all in dict1, but dict1 is more
    for k, v in dict2.items():
        if k not in dict1 or v > dict1[k]:
            return -1
    for k, v in dict1.items():
        if k not in dict2 or v > dict2[k]:
            return 1
    return 0

class Solution:
    def checkInclusion(self, s1: str, s2: str) -> bool:

        left, right = 0, 0 # [left, right) is current window
        window, target = defaultdict(int), defaultdict(int)

        s, t = s2, s1
        for c in s1:
            target[c] += 1
        
        while right < len(s2):
            window[s[right]] += 1
            right += 1

            comp_res = dict1_comp_dict2(window, target)
            # print(left, right, window)
            while left < right and comp_res == 1: # window cover target
                # print("  shrink", left, right, window)
                # shrink window size
                l_char = s[left]
                window[l_char] -= 1
                if window[l_char] == 0:
                    del window[l_char]
                left += 1
                comp_res = dict1_comp_dict2(window, target)
            
            if comp_res == 0:
                return True
        
        return False
        

