class Solution:
    def plusOne(self, digits: List[int]) -> List[int]:
        carry = 1
        for i in range(len(digits)-1, -1, -1):
            tmp = digits[i]
            digits[i] = (digits[i] + carry) % 10
            carry = (tmp + carry) // 10
        if carry == 1:
            digits.insert(0, 1)
        return digits