
class Solution:
    def isPalindrome(self, x: int) -> bool:
        if x < 0:
            return False
        res = 0
        num = x

        while num > 0:
            res = res * 10 + num % 10
            num = num // 10
        
        return x == res