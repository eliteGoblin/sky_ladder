from icecream import ic
import operator

def custom_div(a, b):
    result = a // b
    if a % b and (a < 0) != (b < 0):
        result += 1
    return result

operations = {
    '+': operator.add,
    '-': operator.sub,
    '*': operator.mul,
    '/': custom_div,
}

class Solution:
    def evalRPN(self, tokens: list[str]) -> int:
        operands = []
        for token in tokens:
            if token in ["+", "-", "*", "/"]:
                b = operands.pop()
                a = operands.pop()
                operands.append(operations[token](a, b))
            else:
                operands.append(int(token))
        return operands.pop()
    
s = Solution()

s.evalRPN(["10","6","9","3","+","-11","*","/","*","17","+","5","+"])