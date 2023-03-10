class Solution:
    def mergeAlternately(self, word1: str, word2: str) -> str:
        res = ""
        i, j = 0, 0
        fst = True

        while i < len(word1) and j < len(word2):
            if fst:
                res += word1[i:i+1]
                i += 1
            else:
                res += word2[j:j+1]
                j += 1
            fst = not fst
        
        if i < len(word1):
            res += word1[i:]
        if j < len(word2):
            res += word2[j:]
        
        return res