class Solution:
    def freqAlphabets(self, s: str) -> str:
        res = ""
        
        for i in range(len(s)):
            if s[i:i+1] != "#":
                res += s[i:i+1]
                continue
            num = int(s[i-2:i])
            res = res[:-2]
            res += chr((num - 10) + ord('j'))
        
        print(res)
        final_res = ""
        for c in res:
            if "1" <= c <= "9":
                final_res += chr(ord('a') + int(c) - 1)
            else:
                final_res += c
        
        return final_res
