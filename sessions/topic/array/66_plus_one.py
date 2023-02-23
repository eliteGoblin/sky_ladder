class Solution:
    
def plusOne(self, digits: List[int]) -> List[int]:
    carry = 0
    digits[-1] += 1
    for i in range(len(digits), 0, -1):
        digits[i-1] += carry
        carry = digits[i-1] // 10
        digits[i-1] %= 10
    if carry > 0:
        digits.insert(0, carry)
    return digits