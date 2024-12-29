class Solution:
    def isIsomorphic(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False
        dct = dict()
        for i in range(len(s)):
            if s[i] in dct:
                if dct[s[i]] != t[i]:
                    return False
            else:
                if t[i] in dct.values():
                    return False
                dct[s[i]] = t[i]
        return True