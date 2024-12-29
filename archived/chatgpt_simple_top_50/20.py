class Solution:
    def isValid(self, s: str) -> bool:
        stk = []
        dct = {
            "]": "[",
            "}": "{",
            ")": "(",
        }
        for e in s:
            if e in dct.values():
                stk.append(e)
            else:
                if len(stk) == 0:
                    return False
                if dct[e] != stk.pop():
                    return False
        return len(stk) == 0