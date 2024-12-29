class Solution:
    def plusOne(self, digits: List[int]) -> List[int]:
        carry = 1
        for i in range(-1, -len(digits) - 1, -1):
            if carry == 0:
                break
            digits[i] += carry
            carry = digits[i] // 10
            digits[i] = digits[i] % 10

        if carry > 0:
            msb = 1
            digits[0] = digits[0] % 10
            digits.insert(0, msb)
        
        return digits
