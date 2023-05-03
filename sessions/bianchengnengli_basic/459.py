
def is_repeated(sstr, s) -> bool:
    for i in range(0, len(s), len(sstr)):
        if sstr != s[i:i+sstr]:
            return False
    return True

class Solution:
    def repeatedSubstringPattern(self, s: str) -> bool:
        sub_str_len = 1
        for sub_str_len in range(1, len(s)):
            if len(s) % sub_str_len != 0:
                continue
            if is_repeated(s[:sub_str_len], s):
                return True
        return False