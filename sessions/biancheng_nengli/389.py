class Solution:
    def findTheDifference(self, s: str, t: str) -> str:
        dt = {}
        for c in s:
            if c not in dt:
                dt[c] = 1
            else:
                dt[c] += 1
        
        for c in t:
            if c in dt:
                dt[c] -= 1
                if dt[c] == 0:
                    del dt[c]
            else:
                return c
        
        return ""