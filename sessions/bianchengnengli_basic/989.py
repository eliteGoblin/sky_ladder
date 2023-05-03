class Solution:
    def addToArrayForm(self, num: List[int], k: int) -> List[int]:

        i = len(num) - 1

        carry = 0
        res = []
        while i >= 0 or k > 0:
            op1 = num[i] if i >= 0 else 0
            op2 = k % 10 if k > 0 else 0

            tmp = op1 + op2 + carry
            res.insert(0, tmp % 10)
            carry = tmp // 10

            i -= 1
            k = k // 10

        if carry > 0:
            res.insert(0, carry)
        
        return res