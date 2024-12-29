# use or logic seems better
class Solution:
    def addBinary(self, a: str, b: str) -> str:
        if a == "0":
            return b
        if b == "0":
            return a
        
        i, j = len(a) - 1, len(b) - 1
        res = ""
        carry = 0

        while i >= 0 or j >= 0:
            op1 = int(a[i:i+1]) if i >= 0 else 0
            op2 = int(b[j:j+1]) if j >= 0 else 0

            tmp = op1 + op2 + carry
            carry = tmp // 2
            res = str(tmp % 2) + res

            i -= 1
            j -= 1
        
        if carry == 1:
            res = "1" + res
        
        return res



