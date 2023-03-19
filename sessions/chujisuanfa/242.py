class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        dt = {}
        for c in s:
            if c not in dt:
                dt[c] = 0
            dt[c] += 1
        for c in t:
            if c not in dt:
                return False
            dt[c] -= 1
            if dt[c] == 0:
                del dt[c]
        
        return len(dt) == 0