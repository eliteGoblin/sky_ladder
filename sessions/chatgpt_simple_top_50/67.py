class Solution:
    def addBinary(self, a: str, b: str) -> str:
        a_pos, b_pos = len(a) - 1, len(b) - 1
        carry = 0

        res = ""
        while a_pos >= 0 or b_pos >= 0:
            op1 = int(a) if a_pos >= 0 else 0
            op2 = int(b) if b_pos >= 0 else 0
            tmp = (op1 + op2 + carry)
            carry = tmp // 2
            res = str(tmp % 2) + res
            a_pos -= 1
            b_pos -= 1
        if carry > 0:
            res = str(carry) + res
        return res
