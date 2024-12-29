class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        if len(strs) == 0:
            return 0
        min_length = min(len(s) for s in strs)

        res = ""
        stop = False

        for j in range(0, min_length):
            if stop:
                break
            for s in strs:
                if s[j] != strs[0][j]:
                    stop = True
                    break
            if not stop:
                res += strs[0][j]
        
        return res
