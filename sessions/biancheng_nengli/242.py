class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        dt = {}
        for e in s:
            if e not in dt:
                dt[e] = 1
            else:
                dt[e] += 1
        
        for e in t:
            if e not in dt:
                return False
            if e in dt:
                dt[e] -= 1
                if dt[e] == 0:
                    del dt[e]
        
        return len(dt) == 0