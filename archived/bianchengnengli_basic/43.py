
class Solution:
    def multiply(self, num1: str, num2: str) -> str:
        if num1 == "0" or num2 == "0":
            return "0"
        
        res = 0

        for i in range(len(num2)-1, -1, -1):
            op2 = num2[i:i+1]
            tmp = 0
            for j in range(len(num1)-1, -1, -1):
                tmp += int(num1[j:j+1]) * int(op2) * (10 ** (len(num1) - 1 - j))
            res += tmp * (10 ** (len(num2) - 1 - i))

        return str(res)
