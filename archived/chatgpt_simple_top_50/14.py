class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        res = ""
        for j in range(200):
            for i in range(len(strs)):
                if j >= len(strs[i]):
                    return res
                if strs[i][j] != strs[0][j]:
                    return res
            res += strs[0][j]
        return res
