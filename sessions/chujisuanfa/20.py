
dt = {
    "}": "{",
    ")": "(",
    "]": "[",
}

class Solution:
    def isValid(self, s: str) -> bool:
        stk = []

        for c in s:
            if c in ["(", "{", "["]:
                stk.append(c)
            else:
                if len(stk) == 0:
                    return False
                if dt[c] != stk.pop():
                    return False
        
        return len(stk) == 0